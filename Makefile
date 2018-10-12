BINNAME=scribe
BINARYDIR=$(GOPATH)/bin/
RELEASEBINDIR=bin

all: test build

test: 
	go test

build: 
	go build -o $(BINNAME) && mv $(BINNAME) $(BINARYDIR)

build-win:
	GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -o $(BINNAME).exe && mv $(BINNAME).exe $(RELEASEBINDIR)/

build-linux:
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o $(BINNAME).linux && mv $(BINNAME).linux $(RELEASEBINDIR)/

# ON OSX
build-osx: 
	go build -o $(BINNAME).osx && mv $(BINNAME).osx $(RELEASEBINDIR)/

create-release: build-win build-linux build-osx


