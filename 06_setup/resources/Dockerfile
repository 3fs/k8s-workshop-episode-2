FROM golang:1.14.3-buster AS builder

WORKDIR /app

COPY ./code .

RUN go build -o example-server main.go

# =============================================================================
# FROM gcr.io/distroless/base:latest AS deployable
FROM ubuntu:20.04 as deployable

COPY --from=builder /app/example-server /usr/bin/example-server

# Sourced from https://github.com/arcadeJHS/HTML5-Space-Invaders
COPY resources/game/index.html /game/index.html

COPY resources/upload.gtpl /templates/upload.gtpl
RUN mkdir /uploadfiles

CMD ["/usr/bin/example-server"]
