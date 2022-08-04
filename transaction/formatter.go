package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
