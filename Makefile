build:
	go build -o ./bin/mohit6b/pow-blockchain-go

run: build
	./bin/mohit6b/pow-blockchain-go

test:
	go test ./...