GetPAN
======

GetPAN is an alternative to Carton.

It [has some differences to Carton](Carton.md), but essentially does 
the same job of installing Perl dependencies from a cpanfile.

## Features

- Supports basic cpanfile syntax
- Builds a full dependency tree
- Downloads CPAN archives to local cache
- Resolves dependency installation order
- Supports [SmartPAN](../smartpan/README.md) for accurate dependency resolution
- Supports multiple CPAN mirrors
- Supports BackPAN for old module versions
- Per-module notest and dependency fixes
- Should work on most platforms (requires cpanm!)
- Carton compatible - run `carton install` after `getpan`

## Getting started

Run GetPAN from the command line to install Perl dependencies:

    getpan

This will use the default CPAN source and will install dependencies
found in `./cpanfile`.

You can provide alternative mirrors to install from:

    getpan -cpan http://your.darkpan:1234

If you have a [SmartPAN](../smartpan/README.md), pass in the URL with `-smart`:

    getpan -smart http://your.smartpan:7050

## Command line options

| Option            | Example                          | Description
| ---------         | -------                          | -----------
| -h                | -h                               | Display usage information
| -backpan          | -backpan http://backpan.perl.org | A BackPAN mirror to use (can be specified multiple times)
| -cpan             | -cpan http://www.cpan.org        | A CPAN mirror to use (can be specified multiple times)
| -cpanfile         | -cpanfile app.cpanfile           | The cpanfile to install from
| -cpus             | -cpus 4                          | Number of CPUs to use
| -loglayout        | -loglayout="[%d] %m"             | A github.com/ian-kent/go-log compatible pattern layout
| -loglevel         | -loglevel=TRACE                  | Set log output level (ERROR, INFO, WARN, DEBUG, TRACE)
| -nevertest        | -nevertest                       | Disables all installation tests
| -nocpan           | -nocpan                          | Disables the default CPAN source
| -nobackpan        | -nobackpan                       | Disables the default BackPAN source
| -noinstall        | -noinstall                       | Skips the installation phase
| -notest           | -notest AnyCache                 | Disables tests for a specific module

## cpanfile support

GetPAN supports basic cpanfile syntax.

### Basic syntax

	# Latest version
    requires 'Module::Name';

    # Minimum version
    requires 'Module::Name', '1.02';
    requires 'Module::Name', '>= 1.02';

    # Exact version
    requires 'Module::Name', '== 1.02';

    # Maximum version
    requires 'Module::Name', '<= 1.02';

### Additional syntax

GetPAN also extends cpanfile syntax to fix missed dependencies in CPAN modules:

    requires 'Broken::Module', '== 1.24'; # REQS: Missing::Dep-3.12; Another::Missing::Dep-1.82


## Licence

Copyright ©‎ 2014, Ian Kent (http://www.iankent.eu).

Released under MIT license, see [LICENSE](LICENSE.md) for details.