package products

import (
	"context"
	"fmt"

	"github.com/popeskul/product-store/pkg/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const LogCollectionName = "products"
const DBName = "store"

type RepositoryProducts struct {
	db *mongo.Client
}

func NewRepoProducts(db *mongo.Client) *RepositoryProducts {
	return &RepositoryProducts{
		db: db,
	}
}

func (r *RepositoryProducts) InsertProduct(ctx context.Context, product *domain.Product) error {
	_, err := r.db.Database(DBName).Collection(LogCollectionName).InsertOne(ctx, product)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryProducts) UpdateByName(ctx context.Context, product *domain.Product) error {
	_, err := r.db.Database(DBName).
		Collection(LogCollectionName).
		UpdateOne(
			ctx,
			bson.M{"name": product.Name},
			bson.M{"$set": product},
		)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryProducts) GetByName(ctx context.Context, name string) (*domain.Product, error) {
	var product domain.Product
	err := r.db.Database(DBName).
		Collection(LogCollectionName).
		FindOne(ctx, bson.M{"name": name}).
		Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *RepositoryProducts) Find(ctx context.Context, paging *domain.ListPagingParams, sorting *domain.ListSortingParams) ([]domain.Product, error) {
	opts := options.Find()
	sortOptions := bson.D{{
		Key:   fmt.Sprintf("%s", sorting.SortBy),
		Value: sorting.SortDirection,
	}}
	opts.SetSort(sortOptions)
	opts.SetLimit(int64(paging.PerPage))
	opts.SetSkip(int64(paging.PerPage * (paging.Page - 1)))

	cur, err := r.db.Database(DBName).Collection(LogCollectionName).Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var products []domain.Product
	for cur.Next(ctx) {
		var p domain.Product
		if err = cur.Decode(&p); err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	if err = cur.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
