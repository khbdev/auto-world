# 1. Builder stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o auto ./main.go

# 2. Final stage
FROM alpine:latest

WORKDIR /app

# Git va bash kerak bo‘ladi, agar skript ichida git ishlatilsa
RUN apk add --no-cache git bash

COPY --from=builder /app/auto .

# Container ishga tushganda Go binary ishga tushadi
CMD ["./auto"]
