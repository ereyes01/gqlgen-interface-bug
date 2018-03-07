package main

import (
	"log"
	"net/http"

	"github.com/ereyes01/gqlgen-interface-bug/shapes"
	"github.com/vektah/gqlgen/handler"
)

func main() {
	http.Handle("/", handler.Playground("Shapes", "/query"))
	http.Handle("/query", handler.GraphQL(shapes.MakeExecutableSchema(new(shapes.ShapeResolver))))

	log.Fatal(http.ListenAndServe(":9090", nil))
}
