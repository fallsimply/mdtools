# mdtools
pure go markdown parser with minimal or none regular expressions

table of contents
- [installation](#installation)
- [licensing](#licensing)
- [local modificiation](#local-modification)
- [why internal?](#why-internal)

## installation
run `go get github.com/fallsimply/mdtools` to get an unmodified version of this library

## licensing
mdtools is avalible under the LGPL which allows this library to be used in commercial and/or closed-source projects as long the original or modified source code is made avalible. \
if you or your employer cannot use LGPL licensed libraries consider contributing to the [dual licensing dicussion](https://github.com/fallsimply/mdtools/discussions/2)

## local modification 
if you want to modify this library run `git clone github.com/fallsimply/mdtools` in the module's directory then add `replace github.com/fallsimply/mdtools => ./mdtools` \
if you release this modified version i'd appreciate it if you would let me know in [this discussion thread](https://github.com/fallsimply/mdtools/discussions/1) and/or make a pull request to bring your changes into the core library

## why internal?
internal packages allow for better orginization and ease of development while keeping the api surface small