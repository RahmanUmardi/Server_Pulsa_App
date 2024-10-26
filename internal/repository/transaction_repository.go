package repository

import (
	"database/sql"
	"fmt"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/shared/custom"
	"time"
)

type transactionRepository struct {
	db *sql.DB
}

type TransactionRepository interface {
	Create(payload entity.Transactions) (entity.Transactions, error)
	GetAll() ([]custom.TransactionsReq, error)
	GetById(id string) (custom.TransactionsReq, error)
	// Update(payload entity.Transactions) (entity.Transactions, error)
	// Delete(id string) error
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(payload entity.Transactions) (entity.Transactions, error) {
	parsedDate, err := time.Parse("02-01-2006", payload.TransactionDate)
	if err != nil {
		return entity.Transactions{}, fmt.Errorf("invalid date format. Please use dd-mm-yyyy format: %v", err)
	}

	tx, err := r.db.Begin()
	if err != nil {
		return entity.Transactions{}, err
	}

	//insert into transactions table
	var transactionId string
	insertTransaction := "INSERT INTO transactions (id_merchant, id_user, customer_name, destination_number, transaction_date) VALUES ($1, $2, $3, $4, $5) RETURNING transaction_id"

	if err := tx.QueryRow(insertTransaction, payload.MerchantId, payload.UserId, payload.CustomerName, payload.DestinationNumber, parsedDate).Scan(&transactionId); err != nil {
		tx.Rollback()
		return entity.Transactions{}, err
	}

	payload.TransactionsId = transactionId

	//insert into transaction detail table
	insertTransactionDetail := "INSERT INTO transaction_detail (transaction_id, id_product, price) VALUES ($1, $2, $3) RETURNING transaction_detail_id"

	for i := range payload.TransactionDetail {
		var transactionDetailId string

		if err := tx.QueryRow(insertTransactionDetail, transactionId, payload.TransactionDetail[i].ProductId, payload.TransactionDetail[i].Price).Scan(&transactionDetailId); err != nil {
			tx.Rollback()
			return entity.Transactions{}, err
		}
		payload.TransactionDetail[i].TransactionDetailId = transactionDetailId
		payload.TransactionDetail[i].TransactionsId = transactionId

		//fetch product price from product table
		var productPrice float64

		if err := r.db.QueryRow("SELECT price FROM mst_product WHERE id_product = $1", payload.TransactionDetail[i].ProductId).Scan(&productPrice); err != nil {
			tx.Rollback()
			return entity.Transactions{}, err
		}
		payload.TransactionDetail[i].Price = productPrice
	}
	// commit transaction
	if err := tx.Commit(); err != nil {
		return entity.Transactions{}, err
	}

	payload.TransactionDate = parsedDate.Format("02-01-2006")
	return payload, nil

}

func (r *transactionRepository) GetAll() ([]custom.TransactionsReq, error) {
	selectQuery := `
		SELECT
			t.transaction_id, t.customer_name, t.destination_number, t.transaction_date,
			u.id_user, u.username, u.role,
			m.id_merchant, m.name_merchant, m.address,
			td.transaction_detail_id, td.transaction_id, p.id_product, p.name_provider, p.nominal, p.price
			
		FROM transactions t
		JOIN mst_user u ON t.id_user = u.id_user
		JOIN mst_merchant m ON t.id_merchant = m.id_merchant
		JOIN transaction_detail td ON t.transaction_id = td.transaction_id
		JOIN mst_product p ON td.id_product = p.id_product
		ORDER BY t.transaction_date DESC`

	rows, err := r.db.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactionMap := make(map[string]*custom.TransactionsReq)

	for rows.Next() {
		var (
			transaction       custom.TransactionsReq
			user              custom.UserRes
			merchant          custom.MerchantRes
			transactionDetail custom.TransactionDetailReq
			product           custom.ProductRes
		)

		if err := rows.Scan(
			&transaction.TransactionsId, &transaction.CustomerName, &transaction.DestinationNumber, &transaction.TransactionDate,
			&user.Id_user, &user.Username, &user.Role,
			&merchant.IdMerchant, &merchant.NameMerchant, &merchant.Address,
			&transactionDetail.TransactionDetailId, &transactionDetail.TransactionsId,
			&product.IdProduct, &product.NameProvider, &product.Nominal, &product.Price,
		); err != nil {
			return nil, err
		}

		transactionDetail.Product = product

		if existingTransaction, ok := transactionMap[transaction.TransactionsId]; ok {
			existingTransaction.TransactionDetail = append(existingTransaction.TransactionDetail, transactionDetail)
		} else {
			transaction.User = user
			transaction.Merchant = merchant
			transaction.TransactionDetail = []custom.TransactionDetailReq{transactionDetail}
			transactionMap[transaction.TransactionsId] = &transaction
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	transactions := make([]custom.TransactionsReq, 0, len(transactionMap))
	for _, transaction := range transactionMap {
		transactions = append(transactions, *transaction)
	}

	return transactions, nil
}

func (r *transactionRepository) GetById(id string) (custom.TransactionsReq, error) {
	selectQuery := `
	SELECT
		t.transaction_id, t.customer_name, t.destination_number, t.transaction_date,
		u.id_user, u.username, u.role,
		m.id_merchant, m.name_merchant, m.address,
		td.transaction_detail_id, p.id_product, p.name_provider, p.nominal, p.price
		
	FROM transactions t
	JOIN mst_user u ON t.id_user = u.id_user
	JOIN mst_merchant m ON t.id_merchant = m.id_merchant
	JOIN transaction_detail td ON t.transaction_id = td.transaction_id
	JOIN mst_product p ON td.id_product = p.id_product
	WHERE t.transaction_id = $1
	`
	rows, err := r.db.Query(selectQuery, id)
	if err != nil {
		return custom.TransactionsReq{}, err
	}
	defer rows.Close()

	var transaction custom.TransactionsReq
	transactionDetailMap := make(map[string]custom.TransactionDetailReq)

	//loop and process the first row after checking
	for rows.Next() {
		var (
			user              custom.UserRes
			merchant          custom.MerchantRes
			transactionDetail custom.TransactionDetailReq
			product           custom.ProductRes
		)
		if err := rows.Scan(
			&transaction.TransactionsId, &transaction.CustomerName, &transaction.DestinationNumber, &transaction.TransactionDate,
			&user.Id_user, &user.Username, &user.Role,
			&merchant.IdMerchant, &merchant.NameMerchant, &merchant.Address,
			&transactionDetail.TransactionDetailId,
			&product.IdProduct, &product.NameProvider, &product.Nominal, &product.Price); err != nil {
			return custom.TransactionsReq{}, err
		}
		transaction.User = user
		transaction.Merchant = merchant
		transactionDetail.Product = product

		//store transaction detail in the map
		transactionDetailMap[transactionDetail.TransactionDetailId] = transactionDetail

		//continue iterating if more row are present
		if !rows.Next() {
			break
		}
	}
	for _, detail := range transactionDetailMap {
		transaction.TransactionDetail = append(transaction.TransactionDetail, detail)
	}
	return transaction, nil
}

// func (r *transactionRepository) Update(payload entity.Transactions) (entity.Transactions, error) {
// 	tx, err := r.db.Begin()
// 	if err != nil {
// 		return entity.Transactions{}, err
// 	}
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	// First verify if merchant and user exist
// 	var exists bool
// 	err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM mst_merchant WHERE id_merchant = $1)", payload.MerchantId).Scan(&exists)
// 	if err != nil {
// 		return entity.Transactions{}, err
// 	}
// 	if !exists {
// 		return entity.Transactions{}, errors.New("merchant not found")
// 	}

// 	err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM mst_user WHERE id_user = $1)", payload.UserId).Scan(&exists)
// 	if err != nil {
// 		return entity.Transactions{}, err
// 	}
// 	if !exists {
// 		return entity.Transactions{}, errors.New("user not found")
// 	}

// 	// Update all fields in the transactions table
// 	updateTransaction := `
// 		UPDATE transactions
// 		SET
// 			id_merchant = $1,
// 			id_user = $2,
// 			customer_name = $3,
// 			destination_number = $4,
// 			transaction_date = $5
// 		WHERE transaction_id = $6`

// 	result, err := tx.Exec(
// 		updateTransaction,
// 		payload.MerchantId,
// 		payload.UserId,
// 		payload.CustomerName,
// 		payload.DestinationNumber,
// 		payload.TransactionDate,
// 		payload.TransactionsId,
// 	)
// 	if err != nil {
// 		return entity.Transactions{}, err
// 	}

// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return entity.Transactions{}, err
// 	}
// 	if rowsAffected == 0 {
// 		return entity.Transactions{}, errors.New("transaction not found")
// 	}

// 	// Delete existing transaction details
// 	_, err = tx.Exec("DELETE FROM transaction_detail WHERE transaction_id = $1", payload.TransactionsId)
// 	if err != nil {
// 		return entity.Transactions{}, err
// 	}

// 	// Validate and insert new transaction details
// 	insertTransactionDetail := `
// 		INSERT INTO transaction_detail (transaction_id, id_product, price)
// 		VALUES ($1, $2, $3)
// 		RETURNING transaction_detail_id`

// 	for i := range payload.TransactionDetail {
// 		// Verify if product exists
// 		err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM mst_product WHERE id_product = $1)",
// 			payload.TransactionDetail[i].ProductId).Scan(&exists)
// 		if err != nil {
// 			return entity.Transactions{}, err
// 		}
// 		if !exists {
// 			return entity.Transactions{}, fmt.Errorf("product with ID %s not found",
// 				payload.TransactionDetail[i].ProductId)
// 		}

// 		var (
// 			transactionDetailId string
// 			productPrice        float64
// 		)

// 		// Fetch current product price
// 		if err := tx.QueryRow(
// 			"SELECT price FROM mst_product WHERE id_product = $1",
// 			payload.TransactionDetail[i].ProductId,
// 		).Scan(&productPrice); err != nil {
// 			return entity.Transactions{}, err
// 		}

// 		if err := tx.QueryRow(
// 			insertTransactionDetail,
// 			payload.TransactionsId,
// 			payload.TransactionDetail[i].ProductId,
// 			productPrice,
// 		).Scan(&transactionDetailId); err != nil {
// 			return entity.Transactions{}, err
// 		}

// 		payload.TransactionDetail[i].TransactionDetailId = transactionDetailId
// 		payload.TransactionDetail[i].TransactionsId = payload.TransactionsId
// 		payload.TransactionDetail[i].Price = productPrice
// 	}

// 	if err := tx.Commit(); err != nil {
// 		return entity.Transactions{}, err
// 	}

// 	return payload, nil
// }

// func (r *transactionRepository) Delete(id string) error {
// 	tx, err := r.db.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	// Delete transaction details first due to foreign key constraint
// 	_, err = tx.Exec("DELETE FROM transaction_detail WHERE transaction_id = $1", id)
// 	if err != nil {
// 		return err
// 	}

// 	// Delete main transaction
// 	result, err := tx.Exec("DELETE FROM transactions WHERE transaction_id = $1", id)
// 	if err != nil {
// 		return err
// 	}

// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	if rowsAffected == 0 {
// 		return errors.New("transaction not found")
// 	}

// 	return tx.Commit()
// }
