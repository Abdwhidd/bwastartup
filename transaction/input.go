package transaction

import "bwastartup/user"

type CampaignTransactionInput struct {
	Id   int `uri:"id" binding:"required"`
	User user.User
}
