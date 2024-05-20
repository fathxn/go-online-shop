package transaction

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go-online-shop/infra/response"
)

type Repository interface {
	TransactionDBRepository
	TransactionRepository
	ProductRepository
}

type TransactionDBRepository interface {
	Begin() (tx *sqlx.Tx, err error)
	Rollback(tx *sqlx.Tx) (err error)
	Commit(tx *sqlx.Tx) (err error)
}

type TransactionRepository interface {
	CreateTransactionWithTx(ctx context.Context, tx *sqlx.Tx, trx Transaction) (err error)
}
type ProductRepository interface {
	GetProductBySku(ctx context.Context, productSKU string) (product Product, err error)
	UpdateProductStockWithTx(ctx context.Context, tx *sqlx.Tx, product Product) (err error)
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

	tx, err := s.repo.Begin()
	if err != nil {
		return
	}
	defer s.repo.Rollback(tx)

	if err = s.repo.CreateTransactionWithTx(ctx, tx, trx); err != nil {
		return
	}

	if err = s.repo.UpdateProductStockWithTx(ctx, tx, myProduct); err != nil {
		return
	}

	if err = s.repo.Commit(tx); err != nil {
		return
	}
	return
}
