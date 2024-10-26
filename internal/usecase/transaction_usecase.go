package usecase

import (
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/repository"
	"server-pulsa-app/internal/shared/custom"
)

type transactionUseCase struct {
	repo repository.TransactionRepository
}

type TransactionUseCase interface {
	Create(payload entity.Transactions) (entity.Transactions, error)
	GetAll() ([]custom.TransactionsReq, error)
	GetById(id string) (custom.TransactionsReq, error)
}

func NewTransactionUseCase(repo repository.TransactionRepository) TransactionUseCase {
	return &transactionUseCase{repo: repo}
}

func (u *transactionUseCase) Create(payload entity.Transactions) (entity.Transactions, error) {
	return u.repo.Create(payload)
}

func (u *transactionUseCase) GetAll() ([]custom.TransactionsReq, error) {
	return u.repo.GetAll()
}

func (u *transactionUseCase) GetById(id string) (custom.TransactionsReq, error) {
	return u.repo.GetById(id)
}
