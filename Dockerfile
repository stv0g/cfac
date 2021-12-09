# syntax=docker/dockerfile:1

FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o /scraper cmd/scraper_amqp/*.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /scraper /

CMD [ "/scraper" ]
