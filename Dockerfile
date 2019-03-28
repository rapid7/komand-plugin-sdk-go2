FROM golang:1.12.1-alpine3.9

ENV BASE=/go/src/github.com/rapid7/komand-plugin-sdk-go2
ADD . /go/src/github.com/rapid7/komand-plugin-sdk-go2

ENV SSL_CERT_FILE=/etc/ssl/certs/ca-certificates.crt \
    SSL_CERT_DIR=/etc/ssl/certs \
    REQUESTS_CA_BUNDLE=/etc/ssl/certs/ca-certificates.crt

WORKDIR /go/src/github.com/rapid7/komand-plugin-sdk-go2

RUN apk add --no-cache make git && \
    make deps && \
    make build && \
    make check && \
    make test && \
    apk del make git
