package usecase

import (
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/repository"
	"server-pulsa-app/internal/shared/custom"
)

type transactionUseCase struct {
	repo repository.TransactionRepository
	log  *logger.Logger
}

type TransactionUseCase interface {
	Create(payload entity.Transactions) (entity.Transactions, error)
	GetAll(userId string) ([]custom.TransactionsReq, error)
	GetById(id string) (custom.TransactionsReq, error)
}

func NewTransactionUseCase(repo repository.TransactionRepository, log *logger.Logger) TransactionUseCase {
	return &transactionUseCase{repo: repo, log: log}
}

func (u *transactionUseCase) Create(payload entity.Transactions) (entity.Transactions, error) {
	u.log.Info("Starting to create a new transaction in the usecase layer", nil)
	return u.repo.Create(payload)
}

func (u *transactionUseCase) GetAll(userId string) ([]custom.TransactionsReq, error) {
	u.log.Info("Starting to get all transactions in the usecase layer", nil)
	return u.repo.GetAll(userId)
}

func (u *transactionUseCase) GetById(id string) (custom.TransactionsReq, error) {
	u.log.Info("Starting to get transaction by id in the usecase layer", nil)
	return u.repo.GetById(id)
}
