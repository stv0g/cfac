# syntax=docker/dockerfile:1

# SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
# SPDX-License-Identifier: Apache-2.0

FROM golang:1.24-alpine AS builder

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

FROM alpine:3.22.0

RUN apk --no-cache add \
    ca-certificates \
    tzdata \
    tesseract-ocr

COPY --from=builder /scraper /
COPY --from=builder /measure /

CMD [ "/scraper" ]
