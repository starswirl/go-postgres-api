FROM golang:1.15.2

ENV GO111MODULE=on

COPY src/ /go/src/app/

WORKDIR /go/src/app


RUN go get -u github.com/cosmtrek/air
RUN go get -u github.com/go-delve/delve/cmd/dlv
# ここあんまり理解してない
CMD air -c .air.toml
