FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY ./resource/go.mod ./resource/go.sum ./
# Adjust local directory structure
COPY ./db ../db

RUN go mod download

COPY ./resource ./

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]
