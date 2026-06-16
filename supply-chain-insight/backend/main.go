package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"

	"supply-chain-insight/backend/internal/handlers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := loadEnvFile(firstExistingPath(".env", "../../.env")); err != nil {
		log.Printf("failed to load .env: %v", err)
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to Supabase: %v", err)
	}
	log.Println("connected to Supabase")

	mux := http.NewServeMux()
	handlers.RegisterUserRoutes(mux, db)
	handlers.RegisterProductRoutes(mux, db)

	port := "8080"

	addr := ":" + port
	log.Printf("backend listening on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}

func firstExistingPath(paths ...string) string {
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	return paths[0]
}

func loadEnvFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}

		key = strings.TrimSpace(key)
		value = strings.Trim(strings.TrimSpace(value), `"'`)
		if key != "" {
			os.Setenv(key, value)
		}
	}

	return scanner.Err()
}
