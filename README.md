# mdtools
pure go markdown parser with minimal or none regular expressions

## licensing
mdtools is avalible under the LGPL which allows this library to be used in commercial and/or closed-source projects as long the original or modified source code is made avalible.

if you release a modified version id appreciate it if you would let me know in a discussion thread or make a pull request to bring your changes into the core library

## embeding
run `go get github.com/fallsimply/mdtools` to get an unmodified version of this library

if you want to modify this library run `git clone github.com/fallsimply/mdtools` in the module's directory then add `replace github.com/fallsimply/mdtools => ./mdtools`

## why internal?
internal packages allow for better orginization and ease of development while keeping the api surface small