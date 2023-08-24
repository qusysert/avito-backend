FROM golang:1.18-alpine as builder
ARG version

RUN apk update && \
    apk add --no-cache make && \
    mkdir -p /app
WORKDIR /app

# for cache go mod depends
COPY go.mod .
COPY go.sum .
RUN go mod download

ENV APP_VERSION=$version
COPY . .
RUN make build