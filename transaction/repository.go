package transaction

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetCampaignByID(campaignID int) ([]Transaction, error)
	GetByUserID(userID int) ([]Transaction, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetCampaignByID(campaignID int) ([]Transaction, error) {

	transactions := []Transaction{}
	err := r.db.Where("campaign_id = ?", campaignID).Preload("User").Order("id DESC").Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil

}

func (r *repository) GetByUserID(userID int) ([]Transaction, error) {

	transactions := []Transaction{}
	err := r.db.Where("user_id = ?", userID).Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Order("id DESC").Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil

}
