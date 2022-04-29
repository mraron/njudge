FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY static/ ./static
COPY cmd/ ./cmd
COPY judge/ ./judge
COPY utils/ ./utils
COPY web/ ./web
COPY templates/ ./templates
COPY main.go ./

RUN go build

RUN apk add --update openssl && rm -rf /var/cache/apk/*
RUN openssl genrsa -out key.pem 4096
RUN openssl rsa -in key.pem -pubout -out key.pub.pem -RSAPublicKey_out