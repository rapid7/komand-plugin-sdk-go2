.PHONY= build run clean deps test

GOWORKDIR=$(GOPATH)/src

build: deps
	go-bindata templates/...
	go build

clean:
	rm -f ./plugin-sdk-go2

run: clean build
	./plugin-sdk-go2 -spec="./plugin.spec.slack.yaml" -package="github.com/komand/testplugins/slack/"

deps:
	go get -u github.com/jteeuwen/go-bindata/...
	go get golang.org/x/tools/cmd/goimports

test: clean build
	go list ./... | grep -v /vendor/ | xargs -P4 -L1 go test -v