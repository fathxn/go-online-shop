package transaction

type Product struct {
	Id    int    `db:"id" json:"id"`
	SKU   string `db:"sku" json:"sku"`
	Name  string `db:"name" json:"name"`
	Stock int    `db:"stock" json:"-"`
	Price int    `db:"price" json:"price"`
}

func (p Product) IsExists() bool {
	return p.Id != 0
}
