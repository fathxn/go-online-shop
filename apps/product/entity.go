package product

import (
	"go-online-shop/infra/response"
	"time"
)

type Product struct {
	Id        int       `db:"id"`
	SKU       string    `db:"sku"`
	Name      string    `db:"name"`
	Stock     int16     `db:"stock"`
	Price     int       `db:"price"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (p Product) Validate() (err error) {
	if err = p.ValidateName(); err != nil {
		return
	}

	if err = p.ValidateStock(); err != nil {
		return
	}

	if err = p.ValidatePrice(); err != nil {
		return
	}

	return
}

func (p Product) ValidateName() (err error) {
	if p.Name == "" {
		return response.ErrProductRequired
	}
	if len(p.Name) < 4 {
		return response.ErrProductInvalid
	}
	return
}

func (p Product) ValidateStock() (err error) {
	if p.Stock <= 0 {
		return response.ErrStockInvalid
	}
	return
}

func (p Product) ValidatePrice() (err error) {
	if p.Price <= 0 {
		return response.ErrPriceInvalid
	}
	return
}
