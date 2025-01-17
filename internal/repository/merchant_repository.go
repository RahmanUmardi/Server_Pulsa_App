package repository

import (
	"database/sql"
	"strings"

	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
)

type MerchantRepository interface {
	Create(payload entity.Merchant) (entity.Merchant, error)
	List() ([]entity.Merchant, error)
	Get(id string) (entity.Merchant, error)
	Update(merchant, newMerchant entity.Merchant) (entity.Merchant, error)
	Delete(id string) error
}

type merchantRepository struct {
	db  *sql.DB
	log *logger.Logger
}

func (m *merchantRepository) Create(payload entity.Merchant) (entity.Merchant, error) {
	m.log.Info("Starting to create a new merchant in the repository layer", nil)

	err := m.db.QueryRow("INSERT INTO mst_merchant (id_user, name_merchant, address, id_product, balance) VALUES ($1, $2, $3, $4, $5) RETURNING id_merchant", payload.IdUser, payload.NameMerchant, payload.Address, payload.IdProduct, 0.0).Scan(&payload.IdMerchant)
	if err != nil {
		m.log.Error("Failed to create the merchant: ", err)
		return entity.Merchant{}, err
	}

	m.log.Info("Merchant has been created successfully: ", payload)
	return payload, nil
}

func (m *merchantRepository) List() ([]entity.Merchant, error) {
	var merchants []entity.Merchant
	var rows *sql.Rows
	var err error

	m.log.Info("Starting to retrive all merchant in the repository layer", nil)

	rows, err = m.db.Query("SELECT id_merchant, id_user, name_merchant, address, id_product, balance FROM mst_merchant")

	if err != nil {
		m.log.Error("Failed to retrive the merchant: ", err)
		return nil, err
	}

	for rows.Next() {
		var merchant entity.Merchant

		m.log.Info("Starting to scan all merchant in the repository layer", nil)
		if err := rows.Scan(&merchant.IdMerchant, &merchant.IdUser, &merchant.NameMerchant, &merchant.Address, &merchant.IdProduct, &merchant.Balance); err != nil {
			m.log.Error("Failed to scan the merchant: ", err)
			return nil, err
		}

		m.log.Info("Starting to add merchant in the repository layer", nil)
		merchants = append(merchants, merchant)
	}

	m.log.Info("Getting all merchant was successfully: ", merchants)
	return merchants, nil
}

func (m *merchantRepository) Get(id string) (entity.Merchant, error) {
	var merchant entity.Merchant

	m.log.Info("Starting to retrive a merchant by id in the repository layer", nil)

	if err := m.db.QueryRow("SELECT id_merchant, id_user, name_merchant, address, id_product, balance FROM mst_merchant WHERE id_merchant = $1", id).Scan(&merchant.IdMerchant, &merchant.IdUser, &merchant.NameMerchant, &merchant.Address, &merchant.IdProduct, &merchant.Balance); err != nil {
		m.log.Error("Failed to retrive the merchant: ", err)
		return entity.Merchant{}, err
	}

	m.log.Info("Getting merchant by id was successfully: ", merchant)
	return merchant, nil
}

func (m *merchantRepository) Update(merchant, payload entity.Merchant) (entity.Merchant, error) {
	m.log.Info("Starting to map merchant and payload in the repository layer", nil)

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

	m.log.Info("Starting to update merchant in the repository layer", nil)

	_, err := m.db.Exec("UPDATE mst_merchant SET id_user = $2, name_merchant = $3, address = $4, id_product = $5 WHERE id_merchant = $1", merchant.IdMerchant, merchant.IdUser, merchant.NameMerchant, merchant.Address, merchant.IdProduct)
	if err != nil {
		m.log.Error("Failed to update the merchant: ", err)
		return entity.Merchant{}, err
	}

	m.log.Info("Merchant has been updated successfully: ", merchant)
	return merchant, nil
}

func (m *merchantRepository) Delete(id string) error {
	m.log.Info("Starting to delete merchant in the repository layer", nil)

	_, err := m.db.Exec("DELETE FROM mst_merchant WHERE id_merchant = $1", id)
	if err != nil {
		m.log.Error("Failed to delete the merchant: ", err)
		return err
	}

	m.log.Info("Merchant has been deleted successfully: ", id)
	return nil
}

func NewMerchantRepository(db *sql.DB, log *logger.Logger) MerchantRepository {
	return &merchantRepository{db: db, log: log}
}
