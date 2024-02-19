FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./pkg ./pkg
COPY ./cmd/ ./cmd
COPY ./internal/ ./internal

RUN go build -o /book-api cmd/book-api/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build /book-api /app/book-api

EXPOSE 4100

CMD ["/app/book-api"]
