package transaction

import (
	"bwastartup/user"
	"time"
)

type Transaction struct {
	Id         int       `json:"id"`
	CampaignId int       `json:"campaign_id"`
	UserId     int       `json:"user_id"`
	Amount     int       `json:"amount"`
	Status     string    `json:"status"`
	Code       string    `json:"code"`
	User       user.User `json:"user"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
