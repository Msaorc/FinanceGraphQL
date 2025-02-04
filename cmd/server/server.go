package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/mattn/go-sqlite3"
	"github.com/msaorc/FinanceGraphQL/graph"
	"github.com/msaorc/FinanceGraphQL/internal/database"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	db, err := sql.Open("sqlite3", "./finance.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	lancamentoDB := database.NewLaunch(db)
	categoriaDB := database.NewCategory(db)
	tipoLancamentoDB := database.NewReleaseType(db)
	nivelNecessidadeDB := database.NewNeedLevel(db)
	formaPagamentoDB := database.NewPaymentMethod(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		LancamentoDB:       lancamentoDB,
		CategoriaDB:        categoriaDB,
		TipoLancamentoDB:   tipoLancamentoDB,
		NivelNecessidadeDB: nivelNecessidadeDB,
		FormaPagamentoDB:   formaPagamentoDB,
	}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
