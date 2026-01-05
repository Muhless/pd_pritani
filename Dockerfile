FROM golang:1.24.2

WORKDIR /Coding/backend/golang/pg_pritani

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

CMD ["./main"]

