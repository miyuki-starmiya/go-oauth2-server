FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY ./auth/go.mod ./auth/go.sum ./
# Adjust local directory structure
COPY ./db ../db

RUN go mod download

COPY ./auth .

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]
