package products

import (
	"context"
	"encoding/csv"
	"net/http"
	"strconv"

	"github.com/popeskul/product-store/internal/repository"
	"github.com/popeskul/product-store/pkg/domain"
)

type ServiceProducts struct {
	repo *repository.Repository
}

func NewProductServices(repo *repository.Repository) *ServiceProducts {
	return &ServiceProducts{
		repo: repo,
	}
}

func (s *ServiceProducts) InsertProduct(ctx context.Context, product *domain.Product) error {
	return s.repo.InsertProduct(ctx, product)
}

func (s *ServiceProducts) UpdateByName(ctx context.Context, product *domain.Product) error {
	return s.repo.UpdateByName(ctx, product)
}

func (s *ServiceProducts) Find(
	ctx context.Context,
	paging *domain.ListPagingParams,
	sorting *domain.ListSortingParams,
) (*domain.ListResponse, error) {
	list, err := s.repo.ProductsRepository.Find(ctx, paging, sorting)
	if err != nil {
		return nil, err
	}

	var products []*domain.Product
	for _, p := range list {
		var product = &domain.Product{}
		product.Name = p.Name
		product.Price = p.Price
		products = append(products, product)
	}

	return &domain.ListResponse{
		Products: products,
	}, nil
}

func (s *ServiceProducts) GetByName(ctx context.Context, name string) (*domain.Product, error) {
	return s.repo.GetByName(ctx, name)
}

func (s *ServiceProducts) Fetch(ctx context.Context, req *domain.FetchRequest) ([]*domain.Product, error) {
	data, err := GetProductsCSVFromURL(ctx, req.Url)
	if err != nil {
		return nil, err
	}

	var products []*domain.Product
	for _, p := range data {
		price, err := strconv.ParseFloat(p[1], 32)
		if err != nil {
			return nil, err
		}

		products = append(products, &domain.Product{
			Name:  p[0],
			Price: float32(price),
		})
	}

	return products, nil
}

func GetProductsCSVFromURL(ctx context.Context, url string) ([][]string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	reader := csv.NewReader(resp.Body)
	reader.Comma = ';'
	return reader.ReadAll()
}
