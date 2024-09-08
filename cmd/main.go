package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/rohithrajasekharan/config-store-server/cmd/api"
	"github.com/rohithrajasekharan/config-store-server/config"
	"github.com/rohithrajasekharan/config-store-server/db"
)

func main() {
	connString := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=require",
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBHost,
		config.Envs.DBName,
	)
	db, err := db.NewPostgresStorage(connString)
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB Successfully connected")
}
