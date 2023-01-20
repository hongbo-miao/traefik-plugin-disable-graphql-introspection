FROM golang:1.17 AS build

WORKDIR /go/src/traefik-plugin-disable-graphql-introspection

COPY go.mod ./
RUN go mod download

COPY . /go/src/traefik-plugin-disable-graphql-introspection
