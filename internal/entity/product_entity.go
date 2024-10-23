package entity

type Product struct {
	IdProduct    string  `db:"id_product" json:"idProduct"`
	NameProvider string  `db:"name_provider" json:"nameProvider"`
	Nominal      float64 `db:"nominal" json:"nominal"`
	Price        float64 `db:"price" json:"price"`
	IdSupliyer   string  `db:"id_supliyer" json:"idSupliyer"`
}
