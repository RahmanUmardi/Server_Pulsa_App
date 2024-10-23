package usecase

import (
	"fmt"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/repository"
)

type MerchantUseCase interface {
	RegisterNewMerchant(payload entity.Merchant) (entity.Merchant, error)
	FindAllMerchant() ([]entity.Merchant, error)
	FindMerchantByID(id string) (entity.Merchant, error)
	UpdateMerchant(payload entity.Merchant) (entity.Merchant, error)
	DeleteMerchant(id string) error
}

type merchantUseCase struct {
	repo repository.MerchantRepository
}

func (m *merchantUseCase) RegisterNewMerchant(payload entity.Merchant) (entity.Merchant, error) {
	return m.repo.Create(payload)
}

func (m *merchantUseCase) FindAllMerchant() ([]entity.Merchant, error) {
	return m.repo.List()
}

func (m *merchantUseCase) FindMerchantByID(id string) (entity.Merchant, error) {
	return m.repo.Get(id)
}

func (m *merchantUseCase) UpdateMerchant(payload entity.Merchant) (entity.Merchant, error) {
	merchant, err := m.repo.Get(payload.IdMerchant)
	if err != nil {
		return entity.Merchant{}, fmt.Errorf("merchant ID of \\%s\\ not found", payload.IdMerchant)
	}

	_, err = m.repo.Update(merchant, payload)
	if err != nil {
		return entity.Merchant{}, fmt.Errorf("merchant ID of \\%s\\ not updated", payload.IdMerchant)
	}
	return m.repo.Get(payload.IdMerchant)
}

func (m *merchantUseCase) DeleteMerchant(id string) error {
	_, err := m.repo.Get(id)
	if err != nil {
		return fmt.Errorf("merchant ID of \\%s\\ not found", id)
	}

	return m.repo.Delete(id)
}

func NewMerchantUseCase(repo repository.MerchantRepository) MerchantUseCase {
	return &merchantUseCase{repo: repo}
}
