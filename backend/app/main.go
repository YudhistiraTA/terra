package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/YudhistiraTA/terra/internal/interface/api/rest"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := 8000
	app := rest.NewEntryPoint()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: app,
	}
	log.Printf("Server running on PORT %d", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("\033[31m%v\033[0m", err)
	}
}
