package repository

import (
	"database/sql"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
)

type ProductRepository interface {
	Create(product entity.Product) (entity.Product, error)
	List() ([]entity.Product, error)
	Get(id string) (entity.Product, error)
	Update(product entity.Product) (entity.Product, error)
	Delete(id string) error
}

type productRepository struct {
	db  *sql.DB
	log *logger.Logger
}

func (p *productRepository) Create(product entity.Product) (entity.Product, error) {
	p.log.Info("Starting to create a new product in the repository layer", nil)

	err := p.db.QueryRow("INSERT INTO mst_product (name_provider, nominal, price, id_supliyer) VALUES ($1, $2, $3, $4) RETURNING id_product", product.NameProvider, product.Nominal, product.Price, product.IdSupliyer).Scan(&product.IdProduct)
	if err != nil {
		p.log.Error("Failed to create the product: ", err)
		return entity.Product{}, err
	}

	p.log.Info("Product has been created successfully: ", product)
	return product, nil

}

func (p *productRepository) Get(id string) (entity.Product, error) {
	var product entity.Product

	p.log.Info("Starting to retrive a product by id in the repository layer", nil)

	err := p.db.QueryRow("SELECT id_product, name_provider, nominal, price, id_supliyer FROM mst_product WHERE id_product = $1", id).Scan(&product.IdProduct, &product.NameProvider, &product.Nominal, &product.Price, &product.IdSupliyer)
	if err != nil {
		p.log.Error("Failed to retrive the product: ", err)
		return entity.Product{}, err
	}

	p.log.Info("Getting user by id was successfully: ", product)
	return product, nil
}

func (p *productRepository) List() ([]entity.Product, error) {
	var products []entity.Product

	p.log.Info("Starting to retrive all product in the repository layer", nil)

	rows, err := p.db.Query("SELECT id_product, name_provider, nominal, price, id_supliyer FROM mst_product")
	if err != nil {
		p.log.Error("Failed to retrive the product: ", err)
		return nil, err
	}

	for rows.Next() {
		var product entity.Product

		p.log.Info("Starting to scan all product in the repository layer", nil)
		err := rows.Scan(&product.IdProduct, &product.NameProvider, &product.Nominal, &product.Price, &product.IdSupliyer)
		if err != nil {
			p.log.Error("Failed to scan the product: ", err)
			return nil, err
		}

		p.log.Info("Starting to add product in the repository layer", nil)
		products = append(products, product)
	}

	p.log.Info("Getting all product was successfully: ", products)
	return products, nil
}

func (p *productRepository) Update(product entity.Product) (entity.Product, error) {
	p.log.Info("Starting to update product in the repository layer", nil)
	// Menggunakan id yang diberikan untuk mengupdate product
	_, err := p.db.Exec("UPDATE mst_product SET name_provider = $1, nominal = $2, price = $3, id_supliyer = $4 WHERE id_product = $5", product.NameProvider, product.Nominal, product.Price, product.IdSupliyer, product.IdProduct)
	if err != nil {
		p.log.Error("Failed to update the product: ", err)
		return entity.Product{}, err
	}

	p.log.Info("Product has been updated successfully: ", product)
	return product, nil
}

func (p *productRepository) Delete(id string) error {
	p.log.Info("Starting to delete product in the repository layer", nil)

	_, err := p.db.Exec("DELETE FROM mst_product WHERE id_product = $1", id)
	if err != nil {
		p.log.Error("Failed to delete the product: ", err)
		return err
	}

	p.log.Info("Product has been deleted successfully: ", id)
	return nil
}

func NewProductRepository(db *sql.DB, log *logger.Logger) ProductRepository {
	return &productRepository{db: db, log: log}
}
