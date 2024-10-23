package entity

import "github.com/google/uuid"

type Product struct {
	IdProduct    uuid.UUID `db:"id_product" json:"id_product"`
	NameProvider string    `db:"name_provider" json:"name_provider"`
	Nominal      float64   `db:"nominal" json:"nominal"`
	Price        float64   `db:"price" json:"price"`
	IdSupliyer   uuid.UUID `db:"id_supliyer" json:"id_supliyer"`
}
