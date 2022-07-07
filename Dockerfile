FROM golang:1.18-alpine

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux

COPY go.* ./
RUN go mod download

COPY . .
RUN go build -o ./bin/shopper ./cmd/shopper/shopper.go

EXPOSE 8080

CMD ["./bin/shopper", "-config", "./config/config.yaml"]
