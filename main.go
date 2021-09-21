package main

import (
	"fmt"

	"golang_server_practice/server/gql_server"
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

	go func() {
		gql_server.StartGQLServer()
	}()

	fmt.Println("Server running...")
	select {}
}
