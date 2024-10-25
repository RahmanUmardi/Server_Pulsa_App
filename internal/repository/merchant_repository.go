package repository

import (
	"database/sql"
	"strings"

	"server-pulsa-app/internal/entity"

	"github.com/sirupsen/logrus"
)

// var logMerchant = logger.GetLogger()

type MerchantRepository interface {
	Create(payload entity.Merchant) (entity.Merchant, error)
	List() ([]entity.Merchant, error)
	Get(id string) (entity.Merchant, error)
	Update(merchant, newMerchant entity.Merchant) (entity.Merchant, error)
	Delete(id string) error
}

type merchantRepository struct {
	db *sql.DB
}

func (m *merchantRepository) Create(payload entity.Merchant) (entity.Merchant, error) {
	// logrus.Info("Starting to create a new merchant in the repository layer")

	err := m.db.QueryRow("INSERT INTO mst_merchant (id_user, name_merchant, address, id_product, balance) VALUES ($1, $2, $3, $4, $5) RETURNING id_merchant", payload.IdUser, payload.NameMerchant, payload.Address, payload.IdProduct, payload.Balance).Scan(&payload.IdMerchant)
	if err != nil {
		// logrus.Error("Failed to create the merchant: ", err)
		return entity.Merchant{}, err
	}

	// logrus.Info("Merchant has been created successfully: ", payload)
	return payload, nil
}

func (m *merchantRepository) List() ([]entity.Merchant, error) {
	var merchants []entity.Merchant
	var rows *sql.Rows
	var err error

	// logrus.Info("Starting to retrive all merchant in the repository layer")

	rows, err = m.db.Query("SELECT id_merchant, id_user, name_merchant, address, id_product, balance FROM mst_merchant")

	if err != nil {
		// logrus.Error("Failed to retrive the product: ", err)
		return nil, err
	}

	for rows.Next() {
		var merchant entity.Merchant

		logrus.Info("Starting to scan all merchant in the repository layer")
		if err := rows.Scan(&merchant.IdMerchant, &merchant.IdUser, &merchant.NameMerchant, &merchant.Address, &merchant.IdProduct, &merchant.Balance); err != nil {
			// logrus.Error("Failed to scan the merchant: ", err)
			return nil, err
		}

		// logrus.Info("Starting to add merchant in the repository layer")
		merchants = append(merchants, merchant)
	}

	// logrus.Info("Getting all merchant was successfully: ", merchants)
	return merchants, nil
}

func (m *merchantRepository) Get(id string) (entity.Merchant, error) {
	var merchant entity.Merchant

	// logrus.Info("Starting to retrive a merchant by id in the repository layer")

	if err := m.db.QueryRow("SELECT id_merchant, id_user, name_merchant, address, id_product, balance FROM mst_merchant WHERE id_merchant = $1", id).Scan(&merchant.IdMerchant, &merchant.IdUser, &merchant.NameMerchant, &merchant.Address, &merchant.IdProduct, &merchant.Balance); err != nil {
		logrus.Error("Failed to retrive the merchant: ", err)
		return entity.Merchant{}, err
	}

	// logrus.Info("Getting merchant by id was successfully: ", merchant)
	return merchant, nil
}

func (m *merchantRepository) Update(merchant, payload entity.Merchant) (entity.Merchant, error) {
	// logrus.Info("Starting to map merchant and payload in the repository layer")

	if strings.TrimSpace(payload.IdUser) != "" {
		merchant.IdUser = payload.IdUser
	}
	if strings.TrimSpace(payload.NameMerchant) != "" {
		merchant.NameMerchant = payload.NameMerchant
	}
	if strings.TrimSpace(payload.Address) != "" {
		merchant.Address = payload.Address
	}
	if strings.TrimSpace(payload.IdProduct) != "" {
		merchant.IdProduct = payload.IdProduct
	}
	if payload.Balance != 0 {
		merchant.Balance = payload.Balance
	}

	// logrus.Info("Starting to update merchant in the repository layer")

	_, err := m.db.Exec("UPDATE mst_merchant SET id_user = $2, name_merchant = $3, address = $4, id_product = $5, balance = $6 WHERE id_merchant = $1", merchant.IdMerchant, merchant.IdUser, merchant.NameMerchant, merchant.Address, merchant.IdProduct, merchant.Balance)
	if err != nil {
		// logrus.Error("Failed to update the merchant: ", err)
		return entity.Merchant{}, err
	}

	// logrus.Info("Merchant has been updated successfully: ", merchant)
	return merchant, nil
}

func (m *merchantRepository) Delete(id string) error {
	// logrus.Info("Starting to delete merchant in the repository layer")

	_, err := m.db.Exec("DELETE FROM mst_merchant WHERE id_merchant = $1", id)
	if err != nil {
		// logrus.Error("Failed to delete the merchant: ", err)
		return err
	}

	// logrus.Info("Merchant has been deleted successfully: ", id)
	return nil
}

func NewMerchantRepository(db *sql.DB) MerchantRepository {
	return &merchantRepository{db: db}
}
