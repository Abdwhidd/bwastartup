package transaction

import (
	"bwastartup/campaign"
	"errors"
	"fmt"
)

type Service interface {
	GetTransactionByCampaignId(input CampaignTransactionInput) ([]Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionByCampaignId(input CampaignTransactionInput) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindById(input.Id)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserId != input.User.Id {
		return []Transaction{}, errors.New("User not owner in campaign")
	}

	fmt.Println("INI INPUT USER ID: %d", int(input.User.Id))
	fmt.Println("INI CAMPAIGN USER ID: %d", int(campaign.UserId))

	transaction, err := s.repository.GetCampaignById(input.Id)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
