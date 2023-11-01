package campaign

import "strings"

type CampaignFormatter struct {
	Id               int    `json:"id"`
	UserId           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatterCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Id = campaign.Id
	campaignFormatter.UserId = campaign.UserId
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.Slug = campaign.Slug
	campaignFormatter.ImageURL = "" // nilai default string kosong

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatterCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatterCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

type CampaignFormatterDetail struct {
	Id               int                   `json:"id"`
	Name             string                `json:"name"`
	ShortDescription string                `json:"short_description"`
	Description      string                `json:"description"`
	ImageURL         string                `json:"image_url"`
	GoalAmount       int                   `json:"goal_amount"`
	CurrentAmount    int                   `json:"current_amount"`
	UserId           int                   `json:"user_id"`
	Slug             string                `json:"slug"`
	Perks            []string              `json:"perks"`
	User             CampaignUserFormatter `json:"user"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatterCampaignDetail(campaign Campaign) CampaignFormatterDetail {
	campaignFormatterDetail := CampaignFormatterDetail{}
	campaignFormatterDetail.Id = campaign.Id
	campaignFormatterDetail.Name = campaign.Name
	campaignFormatterDetail.ShortDescription = campaign.ShortDescription
	campaignFormatterDetail.Description = campaign.Description
	campaignFormatterDetail.GoalAmount = campaign.GoalAmount
	campaignFormatterDetail.CurrentAmount = campaign.CurrentAmount
	campaignFormatterDetail.UserId = campaign.UserId
	campaignFormatterDetail.Slug = campaign.Slug
	campaignFormatterDetail.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFormatterDetail.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	campaignFormatterDetail.Perks = perks

	user := campaign.User
	campaignUserFormatter := CampaignUserFormatter{}
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.ImageURL = user.AvatarFileName

	campaignFormatterDetail.User = campaignUserFormatter

	return campaignFormatterDetail

}
