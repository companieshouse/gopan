package gopan

import (
	"compress/gzip"
	"github.com/ian-kent/go-log/log"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func CountIndex(indexes map[string]map[string]*Source) (int, int, int, int) {
	var n1, n2, n3, n4 int

	n1 = len(indexes)
	for fname, _ := range indexes {
		for _, idx := range indexes[fname] {
			n2 += len(idx.Authors)
			for _, auth := range idx.Authors {
				n3 += len(auth.Packages)
				for _, pkg := range auth.Packages {
					n4 += len(pkg.Provides)
				}
			}
		}
	}

	return n1, n2, n3, n4
}

func AppendToIndex(index string, source *Source, author *Author, pkg *Package) {
	out, err := os.OpenFile(index, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Error("Error opening index: %s", err.Error())
		return
	}

	out.Write([]byte(source.Name + " [" + source.URL + "]\n"))
	out.Write([]byte(" " + author.Name + " [" + author.URL + "]\n"))
	out.Write([]byte("  " + pkg.Name + " => " + pkg.URL + "\n"))
	for p, pk := range pkg.Provides {
		out.Write([]byte("   " + p + " (" + pk.Version + "): " + pk.File + "\n"))
	}

	out.Close()
}

func RemoveModule(index string, source *Source, author *Author, pkg *Package) {
	out, err := os.OpenFile(index, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Error("Error opening index: %s", err.Error())
		return
	}

	out.Write([]byte(source.Name + " [" + source.URL + "]\n"))
	out.Write([]byte(" " + author.Name + " [" + author.URL + "]\n"))
	out.Write([]byte("  -" + pkg.Name + " => " + pkg.URL + "\n"))

	out.Close()
}

func SaveIndex(index string, indexes map[string]*Source) {
	// TODO append, but needs to know which stuff is new
	//out, err := os.OpenFile(".gopancache/index", os.O_RDWR|os.O_APPEND, 0660)
	out, err := os.Create(index)
	if err != nil {
		log.Error("Error creating index: %s", err.Error())
	}
	for _, source := range indexes {
		out.Write([]byte(source.Name + " [" + source.URL + "]\n"))
		log.Trace(source.Name)
		for _, author := range source.Authors {
			out.Write([]byte(" " + author.Name + " [" + author.URL + "]\n"))
			log.Trace("    %s", author.Name)
			for _, pkg := range author.Packages {
				out.Write([]byte("  " + pkg.Name + " => " + pkg.URL + "\n"))
				log.Trace("        %s => %s", pkg.Name, pkg.URL)
				for p, pk := range pkg.Provides {
					out.Write([]byte("   " + p + " (" + pk.Version + "): " + pk.File + "\n"))
					log.Trace("          %s (%s): %s", p, pk.Version, pk.File)
				}
			}
		}
	}
	out.Close()
}

func LoadIndex(index string) map[string]*Source {
	indexes := make(map[string]*Source)

	log.Info("Loading cached index file %s", index)

	if _, err := os.Stat(index); err != nil {
		log.Error("Cached index file not found")
		return indexes
	}

	var bytes []byte
	if strings.HasSuffix(index, ".gz") {
		fi, err := os.Open(index)
		if err != nil {
			log.Error("Error reading index: %s", err.Error())
			return indexes
		}
		defer fi.Close()
		gz, err := gzip.NewReader(fi)
		if err != nil {
			log.Error("Error creating gzip reader: %s", err.Error())
			return indexes
		}

		bytes, err = ioutil.ReadAll(gz)
		if err != nil {
			log.Error("Error reading from gzip: %s", err.Error())
			return indexes
		}
	} else {
		var err error
		bytes, err = ioutil.ReadFile(index)
		if err != nil {
			log.Error("Error reading index: %s", err.Error())
			return indexes
		}
	}

	lines := strings.Split(string(bytes), "\n")
	var csource *Source
	var cauth *Author
	var cpkg *Package
	resrcauth := regexp.MustCompile("^\\s*(.*)\\s\\[(.*)\\]\\s*$")
	repackage := regexp.MustCompile("^\\s*(.*)\\s=>\\s(.*)\\s*$")
	reprovides := regexp.MustCompile("^\\s*(.*)\\s\\((.*)\\):\\s(.*)\\s*$")
	for _, l := range lines {
		log.Trace("Line: %s", l)
		if strings.HasPrefix(l, "   ") {
			// provides
			log.Trace("=> Provides")
			match := reprovides.FindStringSubmatch(l)
			if len(match) > 0 {
				if strings.HasPrefix(match[1], "-") {
					log.Trace("  - Is a removal")
					match[1] = strings.TrimPrefix(match[1], "-")
					if _, ok := cpkg.Provides[match[1]]; ok {
						delete(cpkg.Provides, match[1])
					}
					if len(cpkg.Provides) == 0 {
						delete(cauth.Packages, cpkg.Name)
						cpkg = nil
					}
					if len(cauth.Packages) == 0 {
						delete(csource.Authors, cauth.Name)
						cauth = nil
					}
					if len(csource.Authors) == 0 {
						delete(indexes, csource.Name)
						csource = nil
					}
				} else {
					cpkg.Provides[match[1]] = &PerlPackage{
						Name:    match[1],
						Version: match[2],
						Package: cpkg,
						File:    match[3],
					}
				}
			}
		} else if strings.HasPrefix(l, "  ") {
			// its a package
			log.Trace("=> Package")
			match := repackage.FindStringSubmatch(l)
			if len(match) > 0 {
				if strings.HasPrefix(match[1], "-") {
					log.Trace("  - Is a removal")
					match[1] = strings.TrimPrefix(match[1], "-")
					if _, ok := cauth.Packages[match[1]]; ok {
						delete(cauth.Packages, match[1])
					}
					if len(cauth.Packages) == 0 {
						delete(csource.Authors, cauth.Name)
						cauth = nil
					}
					if len(csource.Authors) == 0 {
						delete(indexes, csource.Name)
						csource = nil
					}
				} else {
					if _, ok := cauth.Packages[match[1]]; ok {
						// we've seen this package before
						log.Trace("Seen this package before: %s", match[1])
						cpkg = cauth.Packages[match[1]]
						continue
					}
					cpkg = &Package{
						Name:     match[1],
						URL:      match[2],
						Author:   cauth,
						Provides: make(map[string]*PerlPackage),
					}
					cauth.Packages[match[1]] = cpkg
				}
			}
		} else if strings.HasPrefix(l, " ") {
			// its an author
			log.Trace("=> Author")
			match := resrcauth.FindStringSubmatch(l)
			if len(match) > 0 {
				if strings.HasPrefix(match[1], "-") {
					log.Trace("  - Is a removal")
					match[1] = strings.TrimPrefix(match[1], "-")
					if _, ok := csource.Authors[match[1]]; ok {
						delete(csource.Authors, match[1])
					}
					if len(csource.Authors) == 0 {
						delete(indexes, csource.Name)
						csource = nil
					}
				} else {
					if _, ok := csource.Authors[match[1]]; ok {
						// we've seen this author before
						log.Trace("Seen this author before: %s", match[1])
						cauth = csource.Authors[match[1]]
						continue
					}
					cauth = &Author{
						Name:     match[1],
						URL:      match[2],
						Source:   csource,
						Packages: make(map[string]*Package, 0),
					}
					csource.Authors[match[1]] = cauth
				}
			}
		} else {
			// its a source
			log.Trace("=> Source")
			match := resrcauth.FindStringSubmatch(l)
			if len(match) > 0 {
				if strings.HasPrefix(match[1], "-") {
					log.Trace("  - Is a removal")
					match[1] = strings.TrimPrefix(match[1], "-")
					if _, ok := indexes[match[1]]; ok {
						delete(indexes, match[1])
					}
				} else {
					seen := false
					for _, idx := range indexes {
						if idx.Name == match[1] {
							// we've seen this source before
							log.Trace("Seen this source before: %s", idx.Name)
							csource = idx
							seen = true
							break
						}
					}
					if seen {
						continue
					}
					csource = &Source{
						Name:    match[1],
						URL:     match[2],
						Authors: make(map[string]*Author, 0),
					}
					indexes[csource.Name] = csource
				}
			}
		}
	}

	for _, source := range indexes {
		log.Trace(source.Name)
		for _, author := range source.Authors {
			log.Trace("    %s", author.Name)
			for _, pkg := range author.Packages {
				log.Trace("        %s => %s", pkg.Name, pkg.URL)
			}
		}
	}

	return indexes
}
