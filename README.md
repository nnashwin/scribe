# scribe

[![Build Status](https://travis-ci.org/ru-lai/scribe.svg?branch=master)](https://travis-ci.org/ru-lai/scribe)

> Create and store links in an intuitive way

## Install
1. [Install Golang](https://golang.org/doc/install)
2. [Set up your GOPATH and GOBIN variables](https://github.com/golang/go/wiki/SettingGOPATH)
3. [Add your GOPATH/bin to your path](https://codevenue.wordpress.com/2015/07/26/golang-setting-up-go-development-environment/)
4. Install the package
```
$ go get github.com/ru-lai/scribe
```

5. change your dir to src/github.com/ru-lai/scribe
```
$ cd $GOPATH/src/github.com/ru-lai/scribe
```
and
```
$ go install
```

You should now be able to run the gogit command from the command line.

## Usage
```
NAME:
   scribe - Quick and easy storage / retrieval of links from keywords

USAGE:
   scribe [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR:
   Tyler Boright <ru.lai.development@gmail.com>

COMMANDS:
     addLink, al
      - adds a link to your link repository by clue;
        Example: scribe addLink search www.google.com
          //=> Adds www.google.com to your directory of links under the clue 'search'
     deleteLink, dl
      - deletes a previously defined link by clue;
        Example: scribe deleteLink goog
          //=> Deleted the link to 'google.com' from your link directory
     getLink, gl
      - retrieves a previously defined link by clue and pastes it to your clipboard;
        Example: scribe getLink search
          //=> Pastes www.google.com to your clipboard
     listLinks, ll
      - displays all of your stored clues and links;
        Example: scribe listLinks
          //=> Printing out your links:
            - Link: tyler.com, Clue: cookies
            - Link: google.com, Clue: goog

     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
