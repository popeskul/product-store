package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/popeskul/product-store/internal/config"
	"github.com/popeskul/product-store/internal/db"
	"github.com/popeskul/product-store/internal/handler"
	"github.com/popeskul/product-store/internal/repository"
	"github.com/popeskul/product-store/internal/server"
	"github.com/popeskul/product-store/internal/services"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := db.NewMongoDB(ctx, cfg)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	repo := repository.NewRepo(db)
	service, err := services.NewServices(repo)
	handlers := handler.NewHandler(service)
	srv := server.NewServer(handlers, server.LoggingUnaryInterceptor)

	log.Fatal(srv.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port)))
}
