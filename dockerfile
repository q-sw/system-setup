FROM golang:1.24.6-alpine AS builder

WORKDIR /app
COPY . ./
RUN ls -la
RUN go mod tidy &&\
    go build -o bin/config .

FROM ubuntu:24.04 AS exec
COPY --from=builder /app/bin/config /usr/bin/config
RUN config
