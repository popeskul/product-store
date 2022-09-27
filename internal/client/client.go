package client

import (
	"context"
	"fmt"

	"github.com/popeskul/product-store/pkg/domain"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn       *grpc.ClientConn
	grpcClient domain.ProductsServiceClient
}

func NewClient(host string, port int) (*Client, error) {
	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("%s:%d", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}

	cc := domain.NewProductsServiceClient(conn)

	return &Client{
		conn:       conn,
		grpcClient: cc,
	}, nil
}

func (c *Client) CloseConnection() error {
	return c.conn.Close()
}

func (c *Client) List(ctx context.Context, req *domain.ListRequest) (*domain.ListResponse, error) {
	return c.grpcClient.List(ctx, req)
}

func (c *Client) Fetch(ctx context.Context, req *domain.FetchRequest) (*domain.Empty, error) {
	return c.grpcClient.Fetch(ctx, req)
}
