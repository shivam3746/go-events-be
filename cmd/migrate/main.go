// package main

// import (
// 	"database/sql"
// 	"log"
// 	"os"

// 	"github.com/golang-migrate/migrate"
// 	"github.com/golang-migrate/migrate/database/sqlite3"
// 	"github.com/golang-migrate/migrate/source/file"
// )

// func main() {
// 	if len(os.Args) < 2 {
// 		log.Fatal("Please provide a migration direction: 'up' or 'down'")
// 	}

// 	direction := os.Args[1]

// 	db, err := sql.Open("sqlite3", "./data.db")

// 	if err != nil {
// 		log.Fatalf("Failed to connect to the database: %v", err)
// 	}

// 	defer db.Close()

// 	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})

// 	if err != nil {
// 		log.Fatalf("Failed to create migration instance: %v", err)
// 	}

// 	fSrc, err := (&file.File{}).Open("cmd/migrate/migrations")

// 	if err != nil {
// 		log.Fatalf("Failed to open migration files: %v", err)
// 	}

// 	m, err := migrate.NewWithInstance("file", fSrc, "sqlite3", instance)

// 	if err != nil {
// 		log.Fatalf("Failed to create migration instance: %v", err)
// 	}

// 	switch direction {
// 	case "up":
// 		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
// 			log.Fatalf("Failed to apply migrations: %v", err)
// 		}
// 	case "down":
// 		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
// 			log.Fatalf("Failed to revert migrations: %v", err)
// 		}
// 	default:
// 	}
// }

package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/file"
	_ "modernc.org/sqlite" // Pure Go SQLite driver
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a migration direction: 'up' or 'down'")
	}

	direction := os.Args[1]

	// Use modernc.org/sqlite instead of mattn/go-sqlite3
	db, err := sql.Open("sqlite", "./data.db")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Use the sqlite driver (not sqlite3) from golang-migrate
	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		log.Fatalf("Failed to create migration driver: %v", err)
	}

	fSrc, err := (&file.File{}).Open("cmd/migrate/migrations")
	if err != nil {
		log.Fatalf("Failed to open migration files: %v", err)
	}

	m, err := migrate.NewWithInstance("file", fSrc, "sqlite", driver)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}

	switch direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to apply migrations: %v", err)
		}
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to revert migrations: %v", err)
		}
	default:
		log.Fatalf("Invalid migration direction: %s", direction)
	}
}
