# FROM golang:latest AS builder

# ENV CGO_ENABLED=0
# ENV GOOS=linux
# ENV GOARCH=amd64
# WORKDIR /go/src/github.com/edm20627/go_line_reservation
# COPY . .
# CMD ["go", "run", "main.go"]

# 開発用
FROM golang:latest

WORKDIR /go/src/github.com/edm20627/go_line_reservation

RUN apt-get update -qq && apt-get install -y vim git
RUN GO111MODULE=off go get\
      github.com/pilu/fresh\
      github.com/motemen/gore/cmd/gore\
      github.com/mdempsky/gocode\
      github.com/k0kubun/pp\
      golang.org/x/tools/cmd/goimports\
      golang.org/x/lint/golint\
      github.com/motemen/gobump/cmd/gobump\
      github.com/Songmu/make2help/cmd/make2help

EXPOSE 8080

ENTRYPOINT [ "fresh" ]