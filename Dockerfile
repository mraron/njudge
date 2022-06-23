FROM ubuntu:22.04

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

ARG DEBIAN_FRONTEND=noninteractive
ENV TZ=Europe/Budapest
RUN apt-get update && apt-get install -y ca-certificates openssl tzdata cython3 golang pandoc gccgo pypy3 python3-dev libpython3-all-dev g++ gcc build-essential
RUN go mod download

COPY static/ ./static
COPY cmd/ ./cmd
COPY judge/ ./judge
COPY utils/ ./utils
COPY web/ ./web
COPY templates/ ./templates
COPY main.go ./

RUN go build

RUN openssl genrsa -out key.pem 4096
RUN openssl rsa -in key.pem -pubout -out key.pub.pem -RSAPublicKey_out