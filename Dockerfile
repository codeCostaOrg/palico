# Dockerfile to build a hot-reloadable Go app

FROM golang:1.23.0

WORKDIR /app

RUN go install github.com/air-verse/air@v1.52.3

COPY . .

RUN go mod tidy
