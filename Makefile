#!make

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: clean
clean:
	rm -rf ./bin/*

.PHONY: build
build: clean fmt
	mkdir -p ./bin/
	go build -o ./bin/ged ./main.go

.PHONY: run
run: clean fmt vet
	go run main.go

.PHONY: test
test:
	go test -v ./...
