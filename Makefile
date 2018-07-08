BINNAME=scribe
BINARYDIR=$(GOPATH)/bin/

all: test build

test: 
	go test

build: 
	go build -o $(BINNAME) && mv $(BINNAME) $(BINARYDIR)
