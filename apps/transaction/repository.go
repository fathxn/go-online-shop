package transaction

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func (r repository) Begin() (tx *sql.Tx, err error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Rollback(tx *sql.Tx) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Commit(tx *sql.Tx) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) CreateTransactionWithTx(ctx context.Context, tx *sqlx.Tx, trx Transaction) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) GetProductBySku(ctx context.Context, productSKU string) (product Product, err error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) UpdateProductStockWithTx(ctx context.Context, tx *sqlx.Tx, product Product) (err error) {
	//TODO implement me
	panic("implement me")
}

func newRepository(db *sqlx.DB) repository {
	return repository{db: db}
}
