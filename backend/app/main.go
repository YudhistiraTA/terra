package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/YudhistiraTA/terra/internal/infrastructure/db/sqlc"
	"github.com/YudhistiraTA/terra/internal/interface/api/rest"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func loadDb(ctx context.Context) *pgx.Conn {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")
	if username == "" || password == "" || host == "" || port == "" || database == "" {
		log.Fatalf("\033[31m%s\033[0m", "Please provide database configuration")
	}
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database)

	db, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		log.Fatalf("\033[31m%v\033[0m", err)
	}
	return db
}

func main() {
	godotenv.Load()
	ctx := context.Background()
	db := loadDb(ctx)
	defer db.Close(ctx)

	sqlc_client := sqlc.New(db)

	app := rest.NewRestServer(ctx, sqlc_client)

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		log.Fatalf("\033[31m%s\033[0m", "Please provide server port")
	}
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", serverPort),
		Handler: app,
	}
	log.Printf("Server running on PORT %s", serverPort)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("\033[31m%v\033[0m", err)
	}
}
