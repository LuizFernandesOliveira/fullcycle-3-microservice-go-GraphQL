package main

import (
	"database/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/LuizFernandesOliveira/fullcycle-3-microservice-go-GraphQL/graph"
	"github.com/LuizFernandesOliveira/fullcycle-3-microservice-go-GraphQL/graph/generated"
	"github.com/LuizFernandesOliveira/fullcycle-3-microservice-go-GraphQL/internal/database"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatalf("failad to open database %v", err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	courseDb := database.NewCourse(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		CategoryDB: categoryDb,
		CourseDB:   courseDb,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
