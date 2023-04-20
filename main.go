package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
)

//go:embed graphiql/index.html graphiql/output/*
var graphiql embed.FS

func main() {
	srv := handler.New(NewExecutableSchema(Config{Resolvers: &Resolver{}}))

	// debugging
	srv.AroundFields(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
		rc := graphql.GetFieldContext(ctx)
		fieldName := ""
		if rc.Parent != nil && rc.Parent.Object != "" {
			fieldName = fmt.Sprintf("[ %s.%s.%s ]", rc.Parent.Object, rc.Object, rc.Field.Name)
		} else {
			fieldName = fmt.Sprintf("[ %s.%s ]", rc.Object, rc.Field.Name)
		}
		fmt.Println("  ~>", fieldName)
		res, err = next(ctx)
		fmt.Println("  <~", fieldName, "=>", res, err)
		return res, err
	})
	// ~~

	srv.AddTransport(transport.SSE{})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
	})
	srv.Use(extension.Introspection{})
	http.Handle("/graphql", srv)

	if graphiqlDir, err := fs.Sub(fs.FS(graphiql), "graphiql"); err == nil {
		handler := http.FileServer(http.FS(graphiqlDir))
		http.Handle("/debug/", http.StripPrefix("/debug", handler))
	} else {
		fmt.Println("failed to serve graphiql")
	}

	fmt.Println("listening at http://127.0.0.1:4112")
	if err := http.ListenAndServe(":4112", nil); err != nil {
		fmt.Println("error serving http")
		return
	}
}
