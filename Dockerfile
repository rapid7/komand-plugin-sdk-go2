FROM golang:1.9.7

ENV BASE=/go/src/github.com/rapid7/komand-plugin-sdk-go2
ADD . /go/src/github.com/rapid7/komand-plugin-sdk-go2

WORKDIR /go/src/github.com/rapid7/komand-plugin-sdk-go2

RUN make deps && make build && make check && make test
