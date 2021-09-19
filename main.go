package main

import (
	"fmt"
	"golang_server_practice/server/http_server"
)

func main() {
	http_server.StartHttpServer()

	fmt.Println("Server running...")
}
