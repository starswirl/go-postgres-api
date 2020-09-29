FROM golang:1.15.2

ENV GO111MODULE=on

COPY src/ /go/src/app/

WORKDIR /go/src/app

RUN go build
CMD air -c .air.toml
