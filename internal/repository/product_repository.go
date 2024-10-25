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
	Update(id string, product entity.Product) (entity.Product, error)
	Delete(id string) error
}

type productRepository struct {
	db *sql.DB
}

func (p *productRepository) Create(product entity.Product) (entity.Product, error) {
	log := logger.GetLogger()
	log.Infof("Starting to create a new product in the repository layer")

	err := p.db.QueryRow("INSERT INTO mst_product (name_provider, nominal, price, id_supliyer) VALUES ($1, $2, $3, $4) RETURNING id_product", product.NameProvider, product.Nominal, product.Price, product.IdSupliyer).Scan(&product.IdProduct)
	if err != nil {
		log.Error("Failed to create the product: ", err)
		return entity.Product{}, err
	}

	log.Info("Product has been created successfully: ", product)
	return product, nil

}

func (p *productRepository) Get(id string) (entity.Product, error) {
	var product entity.Product

	log := logger.GetLogger()
	log.Info("Starting to retrive a product by id in the repository layer")

	err := p.db.QueryRow("SELECT id_product, name_provider, nominal, price, id_supliyer FROM mst_product WHERE id_product = $1", id).Scan(&product.IdProduct, &product.NameProvider, &product.Nominal, &product.Price, &product.IdSupliyer)
	if err != nil {
		log.Error("Failed to retrive the product: ", err)
		return entity.Product{}, err
	}

	log.Info("Getting user by id was successfully: ", product)
	return product, nil
}

func (p *productRepository) List() ([]entity.Product, error) {
	var products []entity.Product

	log := logger.GetLogger()
	log.Info("Starting to retrive all product in the repository layer")

	rows, err := p.db.Query("SELECT id_product, name_provider, nominal, price, id_supliyer FROM mst_product")
	if err != nil {
		log.Error("Failed to retrive the product: ", err)
		return nil, err
	}

	for rows.Next() {
		var product entity.Product

		log.Info("Starting to scan all product in the repository layer")
		err := rows.Scan(&product.IdProduct, &product.NameProvider, &product.Nominal, &product.Price, &product.IdSupliyer)
		if err != nil {
			log.Error("Failed to scan the product: ", err)
			return nil, err
		}

		log.Info("Starting to add product in the repository layer")
		products = append(products, product)
	}

	log.Info("Getting all product was successfully: ", products)
	return products, nil
}

func (b *productRepository) Update(id string, product entity.Product) (entity.Product, error) {
	log := logger.GetLogger()
	log.Info("Starting to update product in the repository layer")
	// Menggunakan id yang diberikan untuk mengupdate product
	_, err := b.db.Exec("UPDATE mst_product SET name_provider = $1, nominal = $2, price = $3, id_supliyer = $4 WHERE id_product = $5", product.NameProvider, product.Nominal, product.Price, product.IdSupliyer, id)
	if err != nil {
		log.Error("Failed to update the product: ", err)
		return entity.Product{}, err
	}

	// Mengatur id pada product yang dikembalikan
	product.IdProduct = id
	log.Info("Product has been updated successfully: ", product)
	return product, nil
}

func (p *productRepository) Delete(id string) error {
	log := logger.GetLogger()
	log.Info("Starting to delete product in the repository layer")

	_, err := p.db.Exec("DELETE FROM mst_product WHERE id_product = $1", id)
	if err != nil {
		log.Error("Failed to delete the product: ", err)
		return err
	}

	log.Info("Product has been deleted successfully: ", id)
	return nil
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}
