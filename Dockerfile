# syntax=docker/dockerfile:1

FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN apk --no-cache add \
    tesseract-ocr-dev \
    opencv \
    musl-dev \
    gcc g++

RUN go mod download

COPY ./ ./

ENV TAGS=gosseract

RUN go build -tags ${TAGS} -o /scraper ./cmd/scraper_amqp/
RUN go build -tags ${TAGS} -o /measure ./cmd/measure/

FROM alpine:3.20.3

RUN apk --no-cache add \
    ca-certificates \
    tzdata \
    tesseract-ocr

COPY --from=builder /scraper /
COPY --from=builder /measure /

CMD [ "/scraper" ]
