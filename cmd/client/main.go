package main

import (
	"context"
	"fmt"
	"log"

	"github.com/popeskul/product-store/internal/client"
	"github.com/popeskul/product-store/internal/config"
	"github.com/popeskul/product-store/pkg/domain"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	client, err := client.NewClient(cfg.Server.Host, cfg.Server.Port)
	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}
	defer client.CloseConnection()

	product, err := client.List(context.Background(), &domain.ListRequest{
		Paging: &domain.ListPagingParams{
			Page:    1,
			PerPage: 10,
		},
		Sorting: &domain.ListSortingParams{
			SortBy:        domain.ListSortingParams_name,
			SortDirection: 1,
		},
	})
	if err != nil {
		log.Fatalf("Error fetching product: %v", err)
	}

	fmt.Printf("Product: %v\n", product)

	if _, err = client.Fetch(context.Background(), &domain.FetchRequest{
		Url: "http://164.92.251.245:8080/api/v1/products/",
	}); err != nil {
		log.Fatalf("Error fetching product: %v", err)
	}
}
