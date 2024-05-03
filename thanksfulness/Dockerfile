# syntax=docker/dockerfile:1

##
## Build
##
# https://hub.docker.com/layers/library/golang/1.22-alpine/images/sha256-963da5f97ab931c0df6906e8c0ebc7db28c88d013735ae020f9558c3e6cf0580?context=explore
FROM golang:1.22-alpine AS build

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./
RUN go build -o /thanksfulness

##
## Deploy
##
# https://hub.docker.com/_/alpine
FROM alpine

WORKDIR /
COPY --from=build /thanksfulness /thanksfulness
EXPOSE 8080

CMD [ "/thanksfulness" ]
