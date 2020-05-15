clean:
	rm -rf bin/pap*
	rm -rf bin/seed

build: clean
	go build -o ./bin ./...

develop:
	go run ./cmd/paperplane

seed:
	go run ./cmd/seed

run:
	./bin/paperplane

run-build: build run

cc-linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin-linux ./...