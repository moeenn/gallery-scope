package main

import (
	"gallery-scope/pkg/server"
)

const PORT = ":3000"

func main() {
	server := server.New()
	server.Run(PORT)
}
