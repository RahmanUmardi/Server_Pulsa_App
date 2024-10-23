package repository

import (
	"database/sql"
	"log"
	"strings"

	"server-pulsa-app/internal/entity"
)

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
	err := m.db.QueryRow("INSERT INTO mst_merchant (id_user, name_merchant, address, id_product, balance) VALUES ($1, $2, $3, $4, $5) RETURNING id_merchant", payload.IdUser, payload.NameMerchant, payload.Address, payload.IdProduct, payload.Balance).Scan(&payload.IdMerchant)
	if err != nil {
		log.Printf("MerchantRepository.Create: %v \n", err.Error())
		return entity.Merchant{}, err
	}
	return payload, nil
}

func (m *merchantRepository) List() ([]entity.Merchant, error) {
	var merchants []entity.Merchant
	var rows *sql.Rows
	var err error

	rows, err = m.db.Query("SELECT id_merchant, id_user, name_merchant, address, id_product, balance FROM mst_merchant")

	if err != nil {
		log.Printf("MerchantRepository.List: %v \n", err.Error())
		return nil, err
	}

	for rows.Next() {
		var merchant entity.Merchant
		if err := rows.Scan(&merchant.IdMerchant, &merchant.IdUser, &merchant.NameMerchant, &merchant.Address, &merchant.IdProduct, &merchant.Balance); err != nil {
			log.Printf("MerchantRepository.List.Rows.Next(): %v \n", err.Error())
			return nil, err
		}
		merchants = append(merchants, merchant)
	}

	return merchants, nil
}

func (m *merchantRepository) Get(id string) (entity.Merchant, error) {
	var merchant entity.Merchant
	if err := m.db.QueryRow("SELECT id_merchant, id_user, name_merchant, address, id_product, balance FROM mst_merchant WHERE id_merchant = $1", id).Scan(&merchant.IdMerchant, &merchant.IdUser, &merchant.NameMerchant, &merchant.Address, &merchant.IdProduct, &merchant.Balance); err != nil {
		log.Printf("MerchantRepository.Get: %v \n", err.Error())
		return entity.Merchant{}, err
	}
	return merchant, nil
}

func (m *merchantRepository) Update(merchant, payload entity.Merchant) (entity.Merchant, error) {
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

	_, err := m.db.Exec("UPDATE mst_merchant SET id_user = $2, name_merchant = $3, address = $4, id_product = $5, balance = $6 WHERE id_merchant = $1", merchant.IdMerchant, merchant.IdUser, merchant.NameMerchant, merchant.Address, merchant.IdProduct, merchant.Balance)
	if err != nil {
		return entity.Merchant{}, err
	}

	return merchant, nil
}

func (m *merchantRepository) Delete(id string) error {
	_, err := m.db.Exec("DELETE FROM mst_merchant WHERE id_merchant = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func NewMerchantRepository(db *sql.DB) MerchantRepository {
	return &merchantRepository{db: db}
}
