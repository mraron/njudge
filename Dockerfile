FROM ubuntu:22.04

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

ARG DEBIAN_FRONTEND=noninteractive
ENV TZ=Europe/Budapest
RUN apt-get update && apt-get install -y wget ca-certificates openjdk-8-jdk mono-mcs fpc tzdata cython3 golang pandoc gccgo pypy3 python3-dev g++ gcc build-essential
RUN go mod download && go install github.com/go-delve/delve/cmd/dlv@latest

COPY static/ ./static
COPY migrations/ ./migrations
COPY cmd/ ./cmd
COPY pkg/ ./pkg
COPY templates/ ./templates
COPY main.go ./

RUN go build
