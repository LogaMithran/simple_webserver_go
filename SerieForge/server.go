package main

import (
	"SerieForge/connector"
	"SerieForge/graph"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {

	connector.InitializeDBConnection()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	defer func() {
		err := connector.CloseDbConnection()
		if err != nil {
			fmt.Println("Error in closing the database")
		}
	}()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
