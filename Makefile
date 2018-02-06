.PHONY= build run clean deps test image tag

GOWORKDIR=$(GOPATH)/src
VERSION=$(shell git describe --abbrev=0 --tags)
MAJOR_VERSION=$(shell git describe --abbrev=0 --tags | cut -d"." -f1-2)

build:
	go-bindata -pkg sdk templates/...
	go build -o cmd/plugin-sdk-go/plugin-sdk-go github.com/rapid7/komand-plugin-sdk-go2/cmd/plugin-sdk-go

clean:
	rm -f ./github.com/rapid7/komand-plugin-sdk-go2/cmd/plugin-sdk-go/plugin-sdk-go
	
image:
	docker build -t komand/go-plugin-2 .

tag: image
	@echo version is $(VERSION)
	docker tag komand/go-plugin-2 komand/go-plugin-2:$(VERSION)
	docker tag komand/go-plugin-2:$(VERSION) komand/go-plugin-2:$(MAJOR_VERSION)

deps:
	go get -u github.com/jteeuwen/go-bindata/...
	go get golang.org/x/tools/cmd/goimports
	go get -u honnef.co/go/tools/cmd/gosimple
	go get -u honnef.co/go/tools/cmd/staticcheck
	go get -u honnef.co/go/tools/cmd/unused
	go get -u github.com/golang/dep/cmd/dep

test: clean build
	go list ./... | grep -v /vendor/ | xargs -P4 -L1 go test -v

check:
	staticcheck $$(go list ./... | grep -v /vendor/)
	gosimple $$(go list ./... | grep -v /vendor/)
	unused $$(go list ./... | grep -v /vendor/)