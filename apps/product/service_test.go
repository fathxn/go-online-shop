package product

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-online-shop/external/database"
	"go-online-shop/infra/response"
	"go-online-shop/internal/config"
	"log"
	"testing"
)

var svc service

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	repo := newRepository(db)
	svc = newService(repo)
}

func TestCreateProduct_Success(t *testing.T) {
	req := CreateProductRequestPayload{
		Name:  "Baju Baru",
		Stock: 10,
		Price: 10_000,
	}

	err := svc.CreateProduct(context.Background(), req)
	require.Nil(t, err)
}

func TestCreateProduct_Fail(t *testing.T) {
	t.Run("name is required", func(t *testing.T) {
		req := CreateProductRequestPayload{
			Name:  "",
			Stock: 10,
			Price: 10_000,
		}

		err := svc.CreateProduct(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductRequired, err)
	})
}

func TestListProduct_Success(t *testing.T) {
	pagination := ListProductRequestPayload{
		Cursor: 0,
		Size:   10,
	}

	products, err := svc.ListProducts(context.Background(), pagination)
	require.Nil(t, err)
	require.NotNil(t, products)
	log.Printf("%+v", products)
}