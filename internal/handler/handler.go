package handler

import (
	"context"

	"github.com/popeskul/product-store/internal/services"
	"github.com/popeskul/product-store/pkg/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	services *services.Services
	domain.ProductsServiceServer
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) List(ctx context.Context, req *domain.ListRequest) (*domain.ListResponse, error) {
	return h.services.Find(ctx, req.Paging, req.Sorting)
}

func (h *Handler) Fetch(ctx context.Context, req *domain.FetchRequest) (*domain.Empty, error) {
	data, err := h.services.Fetch(ctx, req)
	if err != nil {
		return &domain.Empty{}, err
	}

	for _, row := range data {
		var product *domain.Product

		item, err := h.services.GetByName(ctx, row.Name)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				if err = h.services.InsertProduct(ctx, row); err != nil {
					return &domain.Empty{}, err
				}
			} else {
				return &domain.Empty{}, err
			}
		}

		if item != nil && item.Price != row.Price {
			item.Price = row.Price
			if err = h.services.InsertProduct(ctx, product); err != nil {
				return &domain.Empty{}, err
			}
		}
	}

	return &domain.Empty{}, nil
}
