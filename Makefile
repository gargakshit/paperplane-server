clean:
	rm -rf bin/paperplane
	rm -rf bin/seed
	rm -rf bin/genkeys

build: clean
	go build -o ./bin ./...

develop:
	go run ./cmd/paperplane

seed:
	go run ./cmd/seed

run:
	./bin/paperplane

run-build: build run

genkeys:
	go run ./cmd/genkeys

cc-linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin-linux ./...