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
	UpdateProductStock(ctx context.Context, productId int, newStock uint) (err error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{repo: repo}
}

func (s service) CreateTransaction(ctx context.Context, req CreateTransactionRequestPayload) (err error) {
	myProduct, err := s.repo.GetProductBySku(ctx, req.ProductSKU)
	if err != nil {
		return
	}

	if !myProduct.IsExists() {
		err = response.ErrNotFound
		return
	}

	trx := NewTransactionFromCreateRequest(req)
	trx.FromProduct(myProduct).
		SetGrandTotal().
		SetPlatformFee(1_000)

	if err = trx.Validate(); err != nil {
		return
	}

	if err = trx.ValidateStok(uint8(myProduct.Stock)); err != nil {
		return
	}

	err = s.repo.CreateTransaction(ctx, trx)
	return
}
