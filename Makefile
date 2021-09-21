.PHONY: start
start:
	go run main.go

.PHONY: grpc-gen
grpc-gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative server/grpc_server/grpc.proto

.PHONY: gql-gen
gql-gen:
	go run github.com/99designs/gqlgen generate
