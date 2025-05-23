# syntax=docker/dockerfile:1
FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN mkdir -p /data

VOLUME /data

RUN go mod download

COPY . ./

RUN go build ./main.go

CMD [ "./main" ]