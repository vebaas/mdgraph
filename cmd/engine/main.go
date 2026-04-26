package main

import (
	"context"
	"fmt"
	"log"

	"mdgraph/internal/node"
	"mdgraph/internal/registry"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	fmt.Println("engine starting")

	reg := registry.NewRegistry()
	node.RegisterBaseTypes(reg)

	pool, err := pgxpool.New(context.Background(), "postgres://mdgraph:mdgraph@localhost:5432/mdgraph")
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	defer pool.Close()

	repo := node.NewPostgresRepository(pool, reg)

	fmt.Println("database connected")
	fmt.Println("repository ready", repo)
}
