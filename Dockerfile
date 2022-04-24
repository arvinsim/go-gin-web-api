ARG GO_VERSION=1.18

FROM golang:${GO_VERSION}-alpine AS builder
ENV PORT=8000

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go ./

EXPOSE $PORT

FROM builder as devbuild
ENV BUILD_ENVIRONMENT=development
RUN go build -o bin/go-gin-web-api-docker-development

FROM builder as prodbuild
ENV BUILD_ENVIRONMENT=production
RUN go build -o bin/go-gin-web-api-docker-production
