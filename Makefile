.PHONY= build run clean deps test

GOWORKDIR=$(GOPATH)/src
VERSION=$(shell git describe --abbrev=0 --tags)
MAJOR_VERSION=$(shell git describe --abbrev=0 --tags | cut -d"." -f1-2)

build:
	go-bindata templates/...
	go build

clean:
	rm -f ./plugin-sdk-go2

run: clean build
	./plugin-sdk-go2 -spec="./plugin.spec.example.yaml" -package="github.com/komand/testplugins/example/"

deps:
	go get -u github.com/jteeuwen/go-bindata/...
	go get golang.org/x/tools/cmd/goimports
	go get -u honnef.co/go/tools/cmd/gosimple
	go get -u honnef.co/go/tools/cmd/staticcheck
	go get -u honnef.co/go/tools/cmd/unused
	go get -u github.com/golang/dep/...

test: clean build
	go list ./... | grep -v /vendor/ | xargs -P4 -L1 go test -v

check:
	staticcheck $$(go list ./... | grep -v /vendor/)
	gosimple $$(go list ./... | grep -v /vendor/)
	unused $$(go list ./... | grep -v /vendor/)
