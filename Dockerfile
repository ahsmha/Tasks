FROM golang:1.20-alpine

RUN apk update
RUN apk add --no-cache git curl make gcc g++
ENV GO111MODULE=on

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o /main
