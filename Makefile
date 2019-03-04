.PHONY= build run clean deps test image tag

GOWORKDIR=$(GOPATH)/src
VERSION=$(shell git describe --abbrev=0 --tags)
MINOR_VERSION=$(shell git describe --abbrev=0 --tags | cut -d"." -f1-2)
MAJOR_VERSION=$(shell git describe --abbrev=0 --tags | cut -d"." -f1)

build:
	go-bindata -pkg sdk templates/...
	go build -o cmd/plugin-sdk-go/plugin-sdk-go github.com/rapid7/komand-plugin-sdk-go2/cmd/plugin-sdk-go
	go build -o cmd/spec-parser/spec-parser github.com/rapid7/komand-plugin-sdk-go2/cmd/spec-parser

clean:
	rm -f ./github.com/rapid7/komand-plugin-sdk-go2/cmd/plugin-sdk-go/plugin-sdk-go
	rm -f ./github.com/rapid7/komand-plugin-sdk-go2/cmd/spec-parser/spec-parser

image:
	docker build -t komand/go-plugin-2 .

tag: image
	@echo version is $(VERSION)
	docker tag komand/go-plugin-2 komand/go-plugin-2:$(VERSION)
	docker tag komand/go-plugin-2:$(VERSION) komand/go-plugin-2:$(MAJOR_VERSION)

deps:
	go build -o $(GOPATH)/bin/go-bindata ./vendor/github.com/rapid7/go-bindata/go-bindata
	go get golang.org/x/tools/cmd/goimports
	go get -u honnef.co/go/tools/cmd/gosimple
	go get -u honnef.co/go/tools/cmd/staticcheck
	go get -u honnef.co/go/tools/cmd/unused
	go get -u github.com/golang/dep/cmd/dep

test: clean build
	go test ./... -count=1

check:
	staticcheck -ignore github.com/rapid7/komand-plugin-sdk-go2/bindata.go:ST1005 $$(go list ./... | grep -v /vendor/)
	gosimple $$(go list ./... | grep -v /vendor/)
	unused $$(go list ./... | grep -v /vendor/)

deploy:
	@echo deploying "$(VERSION)", "$(MINOR_VERSION)", "$(MAJOR_VERSION)" to dockerhub
	@echo docker login -u "********" -p "********"
	@docker login -u "$(KOMAND_DOCKER_USERNAME)" -p "$(KOMAND_DOCKER_PASSWORD)"
	docker tag komand/go-plugin-2 komand/go-plugin-2:$(VERSION)
	docker tag komand/go-plugin-2 komand/go-plugin-2:$(MINOR_VERSION)
	docker tag komand/go-plugin-2 komand/go-plugin-2:$(MAJOR_VERSION)
	docker push komand/go-plugin-2
	docker push komand/go-plugin-2:$(VERSION)
	docker push komand/go-plugin-2:$(MINOR_VERSION)
	docker push komand/go-plugin-2:$(MAJOR_VERSION)
