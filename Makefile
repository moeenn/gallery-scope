BINARY = ./bin/gallery-scope
SRC = ./*/**.go

build:
	@echo "Running Tests..."
	@go test ./...
	@go build -o ${BINARY} ${SRC}
	@echo "\nBinary compiled successfully!"


dev:
	go run ${SRC}

start:
	${BINARY}

test:
	go test ./...

clean:
	rm ${BINARY}
	@echo "Build files cleaned"
