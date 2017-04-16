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
	./cmd/plugin-sdk-go/plugin-sdk-go -spec="$(GOWORKDIR)/github.com/komand/plugin-sdk-go2/specs/plugin.spec.example.yaml" -package="github.com/komand/testplugins/example/"

deps:
	go get -u github.com/jteeuwen/go-bindata/...
	go get golang.org/x/tools/cmd/goimports
	go get -u honnef.co/go/tools/cmd/gosimple
	go get -u honnef.co/go/tools/cmd/staticcheck
	go get -u honnef.co/go/tools/cmd/unused
	go get -u github.com/FiloSottile/gvt

test: clean build
	go list ./... | grep -v /vendor/ | xargs -P4 -L1 go test -v

check:
	staticcheck $$(go list ./... | grep -v /vendor/)
	gosimple $$(go list ./... | grep -v /vendor/)
	unused $$(go list ./... | grep -v /vendor/)

run_all: clean build
	# ./cmd/plugin-sdk-go/plugin-sdk-go -spec="$(GOWORKDIR)/github.com/komand/plugin-sdk-go2/specs/plugin.spec.cacador.yaml" -package="github.com/komand/testplugins/cacador/"
	./cmd/plugin-sdk-go/plugin-sdk-go -spec="$(GOWORKDIR)/github.com/komand/plugin-sdk-go2/specs/plugin.spec.example_types.yaml" -package="github.com/komand/testplugins/example_types/"
	./cmd/plugin-sdk-go/plugin-sdk-go -spec="$(GOWORKDIR)/github.com/komand/plugin-sdk-go2/specs/plugin.spec.example.yaml" -package="github.com/komand/testplugins/example/"
	# ./cmd/plugin-sdk-go/plugin-sdk-go -spec="$(GOWORKDIR)/github.com/komand/plugin-sdk-go2/specs/plugin.spec.fullcontact.yaml" -package="github.com/komand/testplugins/fullcontact/"
	# ./cmd/plugin-sdk-go/plugin-sdk-go -spec="$(GOWORKDIR)/github.com/komand/plugin-sdk-go2/specs/plugin.spec.jq.yaml" -package="github.com/komand/testplugins/jq/"
	# ./cmd/plugin-sdk-go/plugin-sdk-go -spec="$(GOWORKDIR)/github.com/komand/plugin-sdk-go2/specs/plugin.spec.json.yaml" -package="github.com/komand/testplugins/json/"
	# ./cmd/plugin-sdk-go/plugin-sdk-go -spec="$(GOWORKDIR)/github.com/komand/plugin-sdk-go2/specs/plugin.spec.nfs.yaml" -package="github.com/komand/testplugins/nfs/"
	# ./cmd/plugin-sdk-go/plugin-sdk-go -spec="$(GOWORKDIR)/github.com/komand/plugin-sdk-go2/specs/plugin.spec.numverify.yaml" -package="github.com/komand/testplugins/numverify/"
	# ./cmd/plugin-sdk-go/plugin-sdk-go -spec="$(GOWORKDIR)/github.com/komand/plugin-sdk-go2/specs/plugin.spec.postgresql.yaml" -package="github.com/komand/testplugins/postgresql/"
	./cmd/plugin-sdk-go/plugin-sdk-go -spec="$(GOWORKDIR)/github.com/komand/plugin-sdk-go2/specs/plugin.spec.slack.yaml" -package="github.com/komand/testplugins/slack/"
	./cmd/plugin-sdk-go/plugin-sdk-go -spec="$(GOWORKDIR)/github.com/komand/plugin-sdk-go2/specs/plugin.spec.uuid.yaml" -package="github.com/komand/testplugins/uuid/"
	# ./cmd/plugin-sdk-go/plugin-sdk-go -spec="$(GOWORKDIR)/github.com/komand/plugin-sdk-go2/specs/plugin.spec.virustotal.yaml" -package="github.com/komand/testplugins/virustotal/"
	./cmd/plugin-sdk-go/plugin-sdk-go -spec="$(GOWORKDIR)/github.com/komand/plugin-sdk-go2/specs/plugin.spec.timers.yaml" -package="github.com/komand/testplugins/timers/"
