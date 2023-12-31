package campaign

import (
	"bwastartup/user"
	"time"
)

type Campaign struct {
	Id               int
	UserId           int
	Name             string
	ShortDescription string
	Description      string
	GoalAmount       int
	CurrentAmount    int
	Perks            string
	BackerCount      int
	Slug             string
	CreatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdateAt         time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	CampaignImages   []CampaignImage
	User             user.User
}

type CampaignImage struct {
	Id         int
	CampaignId int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
