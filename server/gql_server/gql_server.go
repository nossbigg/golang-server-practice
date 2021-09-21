package gql_server

import (
	"golang_server_practice/graph"
	"golang_server_practice/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
)

const defaultPort = "8081"

func StartGQLServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/query", srv)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
