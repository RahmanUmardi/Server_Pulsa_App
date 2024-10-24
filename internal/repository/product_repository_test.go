package repository

import (
	"testing"

	"server-pulsa-app/internal/entity"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}
	defer db.Close()

	repo := NewProductRepository(db)

	product := entity.Product{
		NameProvider: "Provider A",
		Nominal:      10000,
		Price:        12000,
		IdSupliyer:   "Supplier A",
	}

	mock.ExpectQuery("INSERT INTO mst_product").
		WithArgs(product.NameProvider, product.Nominal, product.Price, product.IdSupliyer).
		WillReturnRows(sqlmock.NewRows([]string{"id_product"}).AddRow("1"))

	createdProduct, err := repo.Create(product)
	if err != nil {
		t.Errorf("error was not expected while creating product: %s", err)
	}

	if createdProduct.IdProduct != "1" {
		t.Errorf("expected product ID to be '1', got '%s'", createdProduct.IdProduct)
	}
}

func TestGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}
	defer db.Close()

	repo := NewProductRepository(db)

	productID := "1"
	expectedProduct := entity.Product{
		IdProduct:    productID,
		NameProvider: "Provider A",
		Nominal:      10000,
		Price:        12000,
		IdSupliyer:   "Supplier A",
	}

	mock.ExpectQuery("SELECT id_product, name_provider, nominal, price, id_supliyer FROM mst_product WHERE id_product = ?").
		WithArgs(productID).
		WillReturnRows(sqlmock.NewRows([]string{"id_product", "name_provider", "nominal", "price", "id_supliyer"}).
			AddRow(expectedProduct.IdProduct, expectedProduct.NameProvider, expectedProduct.Nominal, expectedProduct.Price, expectedProduct.IdSupliyer))

	product, err := repo.Get(productID)
	if err != nil {
		t.Errorf("error was not expected while getting product: %s", err)
	}

	if product != expectedProduct {
		t.Errorf("expected product %+v, got %+v", expectedProduct, product)
	}
}

func TestList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}
	defer db.Close()

	repo := NewProductRepository(db)

	mock.ExpectQuery("SELECT id_product, name_provider, nominal, price, id_supliyer FROM mst_product").
		WillReturnRows(sqlmock.NewRows([]string{"id_product", "name_provider", "nominal", "price", "id_supliyer"}).
			AddRow("1", "Provider A", 10000, 12000, "Supplier A").
			AddRow("2", "Provider B", 20000, 22000, "Supplier B"))

	products, err := repo.List()
	if err != nil {
		t.Errorf("error was not expected while listing products: %s", err)
	}

	if len(products) != 2 {
		t.Errorf("expected 2 products, got %d", len(products))
	}
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}
	defer db.Close()

	repo := NewProductRepository(db)

	productID := "1"
	product := entity.Product{
		NameProvider: "Updated Provider",
		Nominal:      15000,
		Price:        17000,
		IdSupliyer:   "Updated Supplier",
	}

	mock.ExpectExec("UPDATE mst_product SET name_provider = ?, nominal = ?, price = ?, id_supliyer = ? WHERE id_product = ?").
		WithArgs(product.NameProvider, product.Nominal, product.Price, product.IdSupliyer, productID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	updatedProduct, err := repo.Update(productID, product)
	if err != nil {
		t.Errorf("error was not expected while updating product: %s", err)
	}

	if updatedProduct.IdProduct != productID {
		t.Errorf("expected product ID to be '%s', got '%s'", productID, updatedProduct.IdProduct)
	}
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}
	defer db.Close()

	repo := NewProductRepository(db)

	productID := "1"

	mock.ExpectExec("DELETE FROM mst_product WHERE id_product = ?").
		WithArgs(productID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Delete(productID)
	if err != nil {
		t.Errorf("error was not expected while deleting product: %s", err)
	}
}
