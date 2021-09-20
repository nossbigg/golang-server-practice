.PHONY: start
start:
	go run main.go

.PHONY: gen-grpc
gen-grpc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative server/grpc_server/grpc.proto
