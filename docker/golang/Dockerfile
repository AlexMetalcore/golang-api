FROM golang:alpine AS builder

RUN apk add --no-cache git build-base

# Delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN mv /go/bin/dlv /usr/local/bin/dlv

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

EXPOSE 9999 2345