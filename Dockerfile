FROM golang:1.18-alpine

RUN apk add build-base

RUN mkdir /app

WORKDIR /app

ADD . /app

RUN GO111MODULE=on

COPY go.mod .

COPY go.sum .

RUN go build cmd/ham_server/main.go

CMD ["./main"]