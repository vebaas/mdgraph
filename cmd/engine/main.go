package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	fmt.Println("engine starting")

	conn, err := pgx.Connect(context.Background(), "postgres://mdgraph:mdgraph@localhost:5432/mdgraph")
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	defer conn.Close(context.Background())
	//
	fmt.Println("database connected")
}
