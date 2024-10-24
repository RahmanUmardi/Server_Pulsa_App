package usecase

import (
	"fmt"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/repository"
)

var logMerchant = logger.GetLogger()

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
	logMerchant.Info("Starting to create a new merchant in the usecase layer")
	return m.repo.Create(payload)
}

func (m *merchantUseCase) FindAllMerchant() ([]entity.Merchant, error) {
	logMerchant.Info("Starting to retrive all merchant in the usecase layer")
	return m.repo.List()
}

func (m *merchantUseCase) FindMerchantByID(id string) (entity.Merchant, error) {
	logMerchant.Info("Starting to retrive a merchant by id in the usecase layer")
	return m.repo.Get(id)
}

func (m *merchantUseCase) UpdateMerchant(payload entity.Merchant) (entity.Merchant, error) {
	logMerchant.Info("Starting to retrive a merchant by id in the usecase layer")

	merchant, err := m.repo.Get(payload.IdMerchant)
	if err != nil {
		logMerchant.Errorf("Merchant ID %s not found: %v", payload.IdMerchant, err)
		return entity.Merchant{}, fmt.Errorf("merchant ID of \\%s\\ not found", payload.IdMerchant)
	}

	logMerchant.Info("Starting to update merchant in the usecase layer")
	_, err = m.repo.Update(merchant, payload)
	if err != nil {
		logMerchant.Error("Failed to update the merchant: ", err)
		return entity.Merchant{}, fmt.Errorf("merchant ID of \\%s\\ not updated", payload.IdMerchant)
	}

	logMerchant.Infof("Merchant ID %s has been updated successfully: ", payload.IdMerchant)
	return m.repo.Get(payload.IdMerchant)
}

func (m *merchantUseCase) DeleteMerchant(id string) error {
	logMerchant.Info("Starting to retrive a merchant by id in the usecase layer")

	_, err := m.repo.Get(id)
	if err != nil {
		logMerchant.Errorf("Merchant ID %s not found: %v", id, err)
		return fmt.Errorf("merchant ID of \\%s\\ not found", id)
	}

	logMerchant.Info("Merchant has been deleted successfully: ", id)
	return m.repo.Delete(id)
}

func NewMerchantUseCase(repo repository.MerchantRepository) MerchantUseCase {
	return &merchantUseCase{repo: repo}
}
