package transaction

import (
	"bwastartup/campaign"
	"errors"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetTransactionsByCampaignID(input GetTransactionsByCampaignIDInput) ([]Transaction, error)
}

func NewTransactionService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionsByCampaignID(input GetTransactionsByCampaignIDInput) ([]Transaction, error) {

	campaign, err := s.campaignRepository.FindById(input.ID)
	if err != nil {
		return nil, err
	}

	if campaign.UserId != input.User.ID {
		return nil, errors.New("you are not authorized to get this campaign")
	}

	transactions, err := s.repository.GetCampaignByID(input.ID)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
