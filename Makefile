clean:
	rm -rf bin/pap*

build: clean
	go build -o ./bin ./...

develop:
	go run ./cmd/paperplane
