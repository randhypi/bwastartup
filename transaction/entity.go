package transaction

import (
	"bwastartup/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       int
	User       user.User
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
