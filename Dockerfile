# stage 1
FROM golang:1.24.2-alpine AS builder

WORKDIR /Coding/backend/golang/pg_pritani

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# stage 2
FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /Coding/backend/golang/pg_pritani/main .

EXPOSE 8080

CMD ["./main"]