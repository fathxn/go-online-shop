package transaction

import (
	"encoding/json"
	"time"
)

type Transaction struct {
	Id           int             `db:"id"`
	Email        string          `db:"email"`
	ProductId    uint            `db:"product_id"`
	ProductPrice uint            `db:"product_price"`
	Amount       uint8           `db:"amount"`
	SubTotal     uint            `db:"sub_total"`
	PlatformFee  uint            `db:"platform_fee"`
	GrandTotal   uint            `db:"grand_total"`
	Status       uint8           `db:"status"`
	ProductJSON  json.RawMessage `db:"product_snapshot"`
	CreatedAt    time.Time       `db:"created_at"`
	UpdatedAt    time.Time       `db:"updated_at"`
}

func (t *Transaction) SetSubTotal() {
	if t.SubTotal == 0 {
		t.SubTotal = t.ProductPrice * uint(t.Amount)
	}
}

func (t *Transaction) SetGrandTotal() {
	if t.GrandTotal == 0 {
		t.SetSubTotal()
		t.GrandTotal = t.SubTotal + t.PlatformFee
	}
}
