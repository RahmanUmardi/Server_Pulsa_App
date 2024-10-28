package usecase

import (
	"fmt"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/repository"
)

type topupUsecase struct {
	repo repository.TopupRepository
}

type TopupUseCase interface {
	CreateTopup(payload entity.TopupRequest) (string, error)
	UpdateAfterPayment(payload entity.TopupRequest) (string, error)
	GetTopupByMerchantId(idMerchant string) ([]entity.TopupRequestDetail, error)
}

func (t *topupUsecase) CreateTopup(payload entity.TopupRequest) (string, error) {
	data, err := t.repo.CreateTopup(payload)
	if err != nil {
		return "", fmt.Errorf("err: %w", err)
	}

	return data, nil
}

func (t *topupUsecase) UpdateAfterPayment(payload entity.TopupRequest) (string, error) {
	err := t.repo.TxTopupUpdateAfterPayment(payload)
	if err != nil {
		return "", fmt.Errorf("failed to update payment: %w", err)
	}

	return payload.Id, nil
}

func (t *topupUsecase) GetTopupByMerchantId(idMerchant string) ([]entity.TopupRequestDetail, error) {
	data, err := t.repo.GetTopupByMerchantId(idMerchant)
	if err != nil {
		return nil, fmt.Errorf("failed to get topup by merchant id: %w", err)
	}

	return data, nil
}

func NewTopupUsecase(repo repository.TopupRepository) TopupUseCase {
	return &topupUsecase{repo: repo}
}
