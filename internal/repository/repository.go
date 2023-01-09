package repository

import (
	"context"

	"github.com/popeskul/product-store/internal/repository/products"
	"github.com/popeskul/product-store/pkg/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductsRepository interface {
	InsertProduct(ctx context.Context, product *domain.Product) error
	UpdateByName(ctx context.Context, product *domain.Product) error
	GetByName(ctx context.Context, name string) (*domain.Product, error)
	Find(ctx context.Context, paging *domain.ListPagingParams, sorting *domain.ListSortingParams) ([]domain.Product, error)
}

type Repository struct {
	ProductsRepository
}

func NewRepo(db *mongo.Client) *Repository {
	return &Repository{
		ProductsRepository: products.NewRepoProducts(db),
	}
}
