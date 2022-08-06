package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type UserTransactionFormatter struct {
	ID        int               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

func NewCampaignTransactionFormatter(campaignTransaction Transaction) CampaignTransactionFormatter {
	return CampaignTransactionFormatter{
		ID:        campaignTransaction.ID,
		Name:      campaignTransaction.User.Name,
		Amount:    campaignTransaction.Amount,
		CreatedAt: campaignTransaction.CreatedAt,
		UpdatedAt: campaignTransaction.UpdatedAt,
	}
}

func NewCampaignTransactionFormatterList(campaignTransactions []Transaction) []CampaignTransactionFormatter {
	if len(campaignTransactions) == 0 {
		return []CampaignTransactionFormatter{}
	}

	var campaignTransactionFormatterList []CampaignTransactionFormatter
	for _, campaignTransaction := range campaignTransactions {
		campaignTransactionFormatterList = append(campaignTransactionFormatterList, NewCampaignTransactionFormatter(campaignTransaction))
	}
	return campaignTransactionFormatterList
}

func NewUserTransactionFormatter(userTransaction Transaction) UserTransactionFormatter {

	campaign := CampaignFormatter{}
	campaign.Name = userTransaction.Campaign.Name

	isPrimary := false
	if len(userTransaction.Campaign.CampaignImages) > 0 {
		for _, campaignImage := range userTransaction.Campaign.CampaignImages {
			if campaignImage.IsPrimary == 1 {
				isPrimary = true
			} else {
				isPrimary = false
			}

			if isPrimary {
				campaign.ImageURL = campaignImage.FileName
				break
			}
		}
	}

	return UserTransactionFormatter{
		ID:        userTransaction.ID,
		Amount:    userTransaction.Amount,
		Status:    userTransaction.Status,
		CreatedAt: userTransaction.CreatedAt,
		Campaign:  campaign,
	}
}

func NewUserTransactionFormatterList(userTransactions []Transaction) []UserTransactionFormatter {
	if len(userTransactions) == 0 {
		return []UserTransactionFormatter{}
	}

	var userTransactionFormatterList []UserTransactionFormatter
	for _, userTransaction := range userTransactions {
		userTransactionFormatterList = append(userTransactionFormatterList, NewUserTransactionFormatter(userTransaction))
	}
	return userTransactionFormatterList
}
