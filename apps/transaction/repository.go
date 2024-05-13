package transaction

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) repository {
	return repository{db: db}
}

func (r repository) CreateTransaction(ctx context.Context, trx Transaction) (err error) {
	query := `INSERT INTO Transactions (
            		email, product_id, product_price
            		,amount, sub_total, platform_fee
            		,grand_total, status, product_json
            		,created_at, updated_at              
			) VALUES (
			    	:email, :product_id, :product_price
            		,:amount, :sub_total, :platform_fee
            		,:grand_total, :status, :product_json
            		,:created_at, :updated_at              
			)
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
}

func (r repository) GetProductBySku(ctx context.Context, productSKU string) (product Product, err error) {
	//TODO implement me
	panic("implement me")
}
