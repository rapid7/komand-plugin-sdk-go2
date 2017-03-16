.PHONY= build run clean

GOWORKDIR=$(GOPATH)/src

build:
	go build

clean:
	rm -f ./plugin-sdk-go2

run: clean build
	./plugin-sdk-go2 -spec="./plugin.spec.uuid.yaml" -package="github.com/komand/testplugins/uuid/"
