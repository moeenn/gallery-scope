BINARY = ./bin/gallery-scope
SRC = ./*/**.go

build:
	go build -o ${BINARY} ${SRC}

run:
	go run ${SRC}

test:
	go test ./...

clean:
	rm ${BINARY}