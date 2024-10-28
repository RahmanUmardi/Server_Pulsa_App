package repository

import (
	"database/sql"
	"fmt"
	"server-pulsa-app/internal/entity"
	"time"
)

type topupRepository struct {
	db *sql.DB
}

type TopupRepository interface {
	CreateTopup(payload entity.TopupRequest) (string, error)
	GetTopupById(tx *sql.Tx, id string) (entity.TopupRequest, error)
	GetTopupByMerchantId(idMerchant string) ([]entity.TopupRequestDetail, error)
	UpdateStatus(tx *sql.Tx, status, idTopup string) error
	UpdatePaymentMethod(tx *sql.Tx, paymentMethod, idTopup string) error
	UpdateBalanceMerchant(tx *sql.Tx, balance int, idMerchant string) error
	UpdateBalanceSupliyer(tx *sql.Tx, balance int, idSupliyer string) error
	TxTopupUpdateAfterPayment(payload entity.TopupRequest) error
}

func (t *topupRepository) CreateTopup(payload entity.TopupRequest) (string, error) {
	payload.CreatedAt = time.Now()

	query := "INSERT INTO tx_topup (id_merchant, id_supliyer, item_name, amount, payment_method, status, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"

	if err := t.db.QueryRow(query, payload.IdMerchant, payload.IdSupliyer, payload.Item_name, payload.Amount, payload.PaymentMethod, payload.Status, payload.CreatedAt).Scan(&payload.Id); err != nil {
		return "", err
	}

	return payload.Id, nil
}

func (t *topupRepository) GetTopupById(tx *sql.Tx, id string) (entity.TopupRequest, error) {
	var payload entity.TopupRequest

	query := "SELECT * FROM tx_topup WHERE id = $1"

	err := tx.QueryRow(query, id).Scan(&payload.Id, &payload.IdMerchant, &payload.IdSupliyer, &payload.Item_name, &payload.Amount, &payload.PaymentMethod, &payload.Status, &payload.CreatedAt)

	if err == sql.ErrNoRows {
		return entity.TopupRequest{}, fmt.Errorf("topup not found")
	} else if err != nil {
		return entity.TopupRequest{}, err
	}

	return payload, nil
}

func (t *topupRepository) GetTopupByMerchantId(idMerchant string) ([]entity.TopupRequestDetail, error) {
	var payload []entity.TopupRequestDetail

	query := "SELECT t.id, t.id_merchant, t.id_supliyer, s.name_supliyer, t.item_name, t.amount, t.payment_method, t.status, t.created_at FROM tx_topup t JOIN mst_supliyer s ON t.id_supliyer = s.id_supliyer WHERE t.id_merchant = $1"

	rows, err := t.db.Query(query, idMerchant)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item entity.TopupRequestDetail
		var supliyer entity.Supliyer
		err := rows.Scan(&item.Id, &item.IdMerchant, &supliyer.IdSupliyer, &supliyer.NameSupliyer, &item.Item_name, &item.Amount, &item.PaymentMethod, &item.Status, &item.CreatedAt)
		if err != nil {
			return nil, err
		}

		item.IdSupliyer = supliyer
		payload = append(payload, item)
	}

	return payload, nil
}

func (t *topupRepository) UpdateStatus(tx *sql.Tx, status, idTopup string) error {
	query := "UPDATE tx_topup SET status = $1 WHERE id = $2"

	if _, err := tx.Exec(query, status, idTopup); err != nil {
		return fmt.Errorf("failed to update status")
	}

	return nil
}

func (t *topupRepository) UpdatePaymentMethod(tx *sql.Tx, paymentMethod, idTopup string) error {
	query := "UPDATE tx_topup SET payment_method = $1 WHERE id = $2"

	if _, err := tx.Exec(query, paymentMethod, idTopup); err != nil {
		return fmt.Errorf("failed to update payment method")
	}

	return nil
}

func (t *topupRepository) UpdateBalanceMerchant(tx *sql.Tx, balance int, idMerchant string) error {
	query := "UPDATE mst_merchant SET balance = balance + $1 WHERE id_merchant = $2"

	if _, err := tx.Exec(query, balance, idMerchant); err != nil {
		return fmt.Errorf("failed to update balance")
	}

	return nil
}

func (t *topupRepository) UpdateBalanceSupliyer(tx *sql.Tx, balance int, idSupliyer string) error {
	query := "UPDATE mst_supplier SET balance = balance - $1 WHERE supplier_id = $2"

	if _, err := tx.Exec(query, balance, idSupliyer); err != nil {
		return fmt.Errorf("failed to update balance")
	}

	return nil
}

func (t *topupRepository) TxTopupUpdateAfterPayment(payload entity.TopupRequest) error {
	tx, err := t.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction")
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	data, err := t.GetTopupById(tx, payload.Id)
	if err != nil {
		return err
	}

	status := "pending"

	if payload.Status == "settlement" {
		status = "paid"
	} else if payload.Status == "cancelled" {
		status = "cancelled"
	}

	err = t.UpdateStatus(tx, status, data.Id)
	if err != nil {
		return err
	}

	err = t.UpdatePaymentMethod(tx, payload.PaymentMethod, data.Id)
	if err != nil {
		return err
	}

	err = t.UpdateBalanceMerchant(tx, data.Amount, data.IdMerchant)
	if err != nil {
		return err
	}

	err = t.UpdateBalanceSupliyer(tx, data.Amount, data.IdSupliyer)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	fmt.Println("Topup transaction committed")

	return nil
}

func NewTopupRepository(db *sql.DB) TopupRepository {
	return &topupRepository{db: db}
}
