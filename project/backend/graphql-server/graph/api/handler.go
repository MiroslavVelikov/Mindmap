package api

import (
	"mindmap-backend/graphql-server/graph"
	"mindmap-backend/graphql-server/graph/utils"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	log "github.com/sirupsen/logrus"
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func ServerHandler() {
	db, err := utils.ConnectToDB()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	resolver := graph.NewResolver()
	server := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: resolver,
	}))

	server.AddTransport(transport.Options{})
	server.AddTransport(transport.GET{})
	server.AddTransport(transport.POST{})

	server.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	server.Use(extension.Introspection{})
	server.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	router := mux.NewRouter()
	router.Handle("/mindmap/api", server)

	err = http.ListenAndServe(":8081", router)
	if err != nil {
		log.Fatal(err)
	}
}
