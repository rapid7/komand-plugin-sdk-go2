FROM golang:1.9.7

ENV BASE=/go/src/github.com/rapid7/komand-plugin-sdk-go2
ADD . /go/src/github.com/rapid7/komand-plugin-sdk-go2

ENV SSL_CERT_FILE /etc/ssl/certs/ca-certificates.crt
ENV SSL_CERT_DIR /etc/ssl/certs
ENV REQUESTS_CA_BUNDLE  /etc/ssl/certs/ca-certificates.crt

WORKDIR /go/src/github.com/rapid7/komand-plugin-sdk-go2

RUN make deps && make build && make check && make test
