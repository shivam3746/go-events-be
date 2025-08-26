package main

import (
	"database/sql"
	"log"
	"rest-api-gin/internal/database"
	"rest-api-gin/internal/env"

	_ "github.com/joho/godotenv/autoload"
	_ "modernc.org/sqlite"
)

type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {
	db, err := sql.Open("sqlite", "./data.db")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	defer db.Close()

	models := database.NewModels(db)
	app := &application{
		port:      env.GetEnvInt("PORT", 4000),
		jwtSecret: env.GetEnvString("JWT_SECRET", "some-secret-key123456"),
		models:    models,
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
