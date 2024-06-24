FROM golang:1.22.4-bookworm

WORKDIR /app

ARG DEBIAN_FRONTEND=noninteractive
ENV TZ=Europe/Budapest
RUN apt-get update && apt-get install -y pandoc

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go install github.com/go-delve/delve/cmd/dlv@latest

COPY static/ ./static
COPY cmd/ ./cmd
COPY pkg/ ./pkg
COPY internal/ ./internal
COPY main.go ./

RUN go build
