# syntax=docker/dockerfile:1

## Build
FROM golang:1.19.2-buster AS builder

WORKDIR /tempest-administration-service

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY config/*.yaml ./

COPY . .
COPY *.go ./

RUN go build -o /tempest-administration-service

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /tempest-administration-service ./

EXPOSE 8080

ENTRYPOINT ["/tempest-administration-service"]