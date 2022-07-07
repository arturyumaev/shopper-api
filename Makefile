build:
	go mod download && go build -o ./bin/shopper ./cmd/shopper/shopper.go

run: build
	./bin/shopper -config ./config/config.yaml
