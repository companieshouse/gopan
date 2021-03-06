package gopan

import (
	"github.com/ian-kent/go-log/log"
	"regexp"
	"strconv"
	"strings"
)

type Package struct {
	Author    *Author
	Name      string
	URL       string
	Provides  map[string]*PerlPackage
	cachedVer float64
}

type PerlPackage struct {
	Package *Package
	Name    string
	Version string
	File    string
}

func (p *Package) String() string {
	return p.Name
}

func (p *Package) VirtualURL() string {
	return p.Author.Source.Name + "/authors/id/" + p.AuthorURL()
}

func (p *Package) AuthorURL() string {
	return p.Author.Name[:1] + "/" + p.Author.Name[:2] + "/" + p.Author.Name + "/" + p.Name
}

var fnToVer = regexp.MustCompile("(.*)-([v_\\-\\.0-9a-zA-Z]*).tar.gz")

func (p *Package) Version() float64 {
	if p.cachedVer > 0 {
		return p.cachedVer
	}

	v := ""

	// try and match against a provided package
	for _, prov := range p.Provides {
		if len(prov.Version) > 0 && prov.Version != "undef" {
			if len(v) == 0 {
				log.Trace("No version cached, using first version found [%s] from [%s]", prov.Version, prov.Name)
				v = prov.Version
			} else {
				if p.Name == strings.Replace(prov.Name, "::", "-", -1)+"-"+prov.Version {
					log.Trace("Version cached but found better match, using [%s] from [%s]", prov.Version, prov.Name)
					v = prov.Version
				}
			}
		}
	}

	if len(v) == 0 {
		matches := fnToVer.FindStringSubmatch(p.Name)
		if len(matches) >= 3 {
			log.Trace("Found regex match: %s", matches[2])
			v = matches[2]
		}
	}

	p.cachedVer = VersionFromString(v)

	return p.cachedVer
}

// An attempt at turning random version numbers into comparable floats!
func VersionFromString(version string) float64 {
	if len(version) == 0 || version == "undef" {
		return float64(0)
	}

	if strings.HasPrefix(version, "v") {
		version = strings.TrimPrefix(version, "v")
	}
	version = strings.Replace(version, "-", ".", -1)
	version = strings.Replace(version, "_", ".", -1)
	parts := strings.Split(version, ".")
	version = parts[0] + "." + strings.Join(parts[1:], "")

	f, _ := strconv.ParseFloat(version, 64)
	return f
}
