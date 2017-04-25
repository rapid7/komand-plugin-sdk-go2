FROM golang:1.8.1

# dependencies 
RUN go get github.com/jteeuwen/go-bindata/...
RUN go get github.com/golang/lint/golint

# gometalinter
RUN go get github.com/alecthomas/gometalinter
RUN gometalinter --install

ENV BASE=/go/src/github.com/komand/plugin-sdk-go2
ADD . /go/src/github.com/komand/plugin-sdk-go2

WORKDIR /go/src/github.com/komand/plugin-sdk-go2
RUN make build && make test