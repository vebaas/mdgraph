package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	dbURL := "postgres://mdgraph:mdgraph@localhost:5432/mdgraph?sslmode=disable"

	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		log.Fatal("failed to create migrator:", err)
	}
	defer m.Close()

	if len(os.Args) < 2 {
		fmt.Println("usage: migrate [up|down]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("migration up failed:", err)
		}
		fmt.Println("migrations applied")

	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("migration down failed:", err)
		}
		fmt.Println("migrations rolled back")

	default:
		fmt.Println("unknown command, use up or down")
		os.Exit(1)
	}
}
