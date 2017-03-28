.PHONY= build run clean deps test

GOWORKDIR=$(GOPATH)/src
VERSION=$(shell git describe --abbrev=0 --tags)
MAJOR_VERSION=$(shell git describe --abbrev=0 --tags | cut -d"." -f1-2)

build:
	go-bindata -pkg sdk templates/...
	go build -o cmd/plugin-sdk-go/plugin-sdk-go github.com/komand/plugin-sdk-go2/cmd/plugin-sdk-go

clean:
	rm -f ./github.com/komand/plugin-sdk-go2/cmd/plugin-sdk-go/plugin-sdk-go

run: clean build
	./cmd/plugin-sdk-go/plugin-sdk-go -spec="$(GOWORKDIR)/github.com/komand/plugin-sdk-go2/specs/plugin.spec.timers.yaml" -package="github.com/komand/testplugins/timers/"

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
