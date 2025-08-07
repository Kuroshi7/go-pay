package main

import (
	"fmt"
	"log"
	"os"
	"database/sql"
	"github.com/kuroshi7/go-pay/internal/repository"
	"github.com/kuroshi7/go-pay/internal/service"
	"github.com/kuroshi7/go-pay/internal/web/server"
	"github.com/joho/godotenv"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {	
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "db"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "password"),
		getEnv("DB_NAME", "go_pay"),
		getEnv("DB_SSLMODE", "disable"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)
	port := getEnv("HTTP_PORT", "8080")
	server := server.NewServer(accountService, port)
	server.ConfigureRoutes()

	server.Start()
	log.Printf("Server running on port %s", port)
	if err := server.Start(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}