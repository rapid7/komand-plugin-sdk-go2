.PHONY= build run clean deps test

GOWORKDIR=$(GOPATH)/src

build:
	go-bindata templates/...
	go build

clean:
	rm -f ./plugin-sdk-go2

run: clean build
	./plugin-sdk-go2 -spec="./plugin.spec.uuid.yaml" -package="github.com/komand/testplugins/uuid/"

deps:
	go get -u github.com/jteeuwen/go-bindata/...

test: clean build
	go list ./... | grep -v /vendor/ | xargs -P4 -L1 go test -v