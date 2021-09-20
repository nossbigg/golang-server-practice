package main

import (
	"fmt"

	"golang_server_practice/server/grpc_server"
	"golang_server_practice/server/http_server"
)

func main() {
	go func() {
		http_server.StartHttpServer()
	}()

	go func() {
		grpc_server.StartGRPCServer()
	}()

	fmt.Println("Server running...")
	select {}
}
