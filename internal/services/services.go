package services

import (
	"context"
	"github.com/popeskul/product-store/internal/repository"
	"github.com/popeskul/product-store/internal/services/products"
	"github.com/popeskul/product-store/pkg/domain"
)

type ProductsService interface {
	InsertProduct(ctx context.Context, product *domain.Product) error
	UpdateByName(ctx context.Context, product *domain.Product) error
	GetByName(ctx context.Context, name string) (*domain.Product, error)
	Find(ctx context.Context, paging *domain.ListPagingParams, sorting *domain.ListSortingParams) (*domain.ListResponse, error)
	Fetch(ctx context.Context, in *domain.FetchRequest) ([]*domain.Product, error)
}

type Services struct {
	ProductsService
}

func NewServices(repo *repository.Repository) (*Services, error) {
	return &Services{
		ProductsService: products.NewProductServices(repo),
	}, nil
}
