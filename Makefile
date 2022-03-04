BINARY = ./bin/gallery-scope
SRC = ./*/**.go

build:
	go build -o ${BINARY} ${SRC}

run:
	go run ${SRC}

clean:
	rm ${BINARY}