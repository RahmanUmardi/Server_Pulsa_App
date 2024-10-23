package repository

import (
	"database/sql"
	"server-pulsa-app/internal/entity"
)

type ProductRepository interface {
	Create(product entity.Product) (entity.Product, error)
	List() ([]entity.Product, error)
	Get(id string) (entity.Product, error)
	Update(id string, product entity.Product) (entity.Product, error)
	Delete(id string) error
}

type productRepository struct {
	db *sql.DB
}

func (p *productRepository) Create(product entity.Product) (entity.Product, error) {
	err := p.db.QueryRow("INSERT INTO mst_product (name_provider, nominal, price, id_supliyer) VALUES ($1, $2, $3, $4) RETURNING id_product", product.NameProvider, product.Nominal, product.Price, product.IdSupliyer).Scan(&product.IdProduct)
	if err != nil {
		return entity.Product{}, err
	}

	return product, nil

}

func (p *productRepository) Get(id string) (entity.Product, error) {
	var product entity.Product
	err := p.db.QueryRow("SELECT id_product, name_provider, nominal, price, id_supliyer FROM mst_product WHERE id_product = $1", id).Scan(&product.IdProduct, &product.NameProvider, &product.Nominal, &product.Price, &product.IdSupliyer)
	if err != nil {
		return entity.Product{}, err
	}

	return product, nil
}

func (p *productRepository) List() ([]entity.Product, error) {
	var products []entity.Product
	rows, err := p.db.Query("SELECT id_product, name_provider, nominal, price, id_supliyer FROM mst_product")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.IdProduct, &product.NameProvider, &product.Nominal, &product.Price, &product.IdSupliyer)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (b *productRepository) Update(id string, product entity.Product) (entity.Product, error) {
	// Menggunakan id yang diberikan untuk mengupdate buku
	_, err := b.db.Exec("UPDATE mst_product SET name_provider = $1, nominal = $2, price = $3, id_supliyer = $4 WHERE id_product = $5", product.NameProvider, product.Nominal, product.Price, product.IdSupliyer, id)
	if err != nil {
		return entity.Product{}, err
	}

	// Mengatur id pada product yang dikembalikan
	product.IdProduct = id
	return product, nil
}

func (p *productRepository) Delete(id string) error {
	_, err := p.db.Exec("DELETE FROM mst_product WHERE id_product = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func NewproductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}
