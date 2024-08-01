# syntax=docker/dockerfile:1

FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o order-api cmd/api/main.go

FROM scratch

COPY --from=builder /app/order-api /order-api

EXPOSE 3001

CMD [ "/order-api" ]