package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "postgres"
	dbPassword = "password"
	dbName     = "python_DB"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("❌ Failed to open database:", err)
	}

	// Connection pool settings
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(1 * time.Hour)

	// Test connection
	if err := db.Ping(); err != nil {
		log.Fatal("❌ Cannot reach DB:", err)
	}
	log.Println("✅ Database connection successful.")

	// Simulate 50 concurrent queries to force open 50 connections
	var wg sync.WaitGroup
	wg.Add(500)

	for i := 0; i < 500; i++ {
		go func(id int) {
			defer wg.Done()
			var now time.Time
			err := db.QueryRow("SELECT NOW()").Scan(&now)
			if err != nil {
				log.Printf("🔁 Goroutine %d failed: %v\n", id, err)
				return
			}
			log.Printf("📦 Goroutine %d → DB time: %s\n", id, now.Format(time.RFC3339))
		}(i)
	}

	wg.Wait()
	log.Println("🎯 All 500 queries executed.")
}
