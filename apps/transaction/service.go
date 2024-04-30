package transaction

import (
	"context"
	"go-online-shop/infra/response"
)

type Repository interface {
	TransactionRepository
	ProductRepository
}
type TransactionRepository interface {
	CreateTransaction(ctx context.Context, trx Transaction) (err error)
}
type ProductRepository interface {
	GetProductBySku(ctx context.Context, productSKU string) (product Product, err error)
}

type service struct {
	repo Repository
}

func (s service) CreateTransaction(ctx context.Context, email, productSKU string) (err error) {
	myProduct, err := s.repo.GetProductBySku(ctx, productSKU)
	if err != nil {
		return
	}

	if myProduct.IsExists() {
		err = response.ErrNotFound
		return
	}

	trx := NewTransaction(email)
	trx.FromProduct(myProduct)
	trx.SetProductJSON(myProduct)

	return
}
