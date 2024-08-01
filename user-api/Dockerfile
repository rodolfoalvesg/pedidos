# syntax=docker/dockerfile:1

FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o user-api cmd/api/main.go

FROM scratch

COPY --from=builder /app/user-api /user-api

EXPOSE 3000

CMD [ "/user-api" ]