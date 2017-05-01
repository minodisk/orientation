FROM golang:1.8.1

RUN go get -u \
      github.com/golang/dep/... \
    dep ensure

WORKDIR /go/src/github.com/minodisk/orientation
COPY . .
