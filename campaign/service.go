package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	FindCampaign(userId int) ([]Campaign, error)
	FindCampaignById(input CampaignDetailInput) (Campaign, error)
	CreateCampaign(input CampaignCreateInput) (Campaign, error)
	UpdateCampaign(inputId CampaignDetailInput, inputData CampaignCreateInput) (Campaign, error)
	SaveCampaignImage(input CreateCampaignImageInput, filelocation string) (CampaignImage, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindCampaign(userId int) ([]Campaign, error) {
	if userId != 0 {
		campaigns, err := s.repository.FindByUserId(userId)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) FindCampaignById(input CampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindById(input.Id)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *service) CreateCampaign(input CampaignCreateInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.User.Id = input.User.Id

	slugCampaign := fmt.Sprintf("%s %d", input.Name, input.User.Id)
	campaign.Slug = slug.Make(slugCampaign)

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil
}

func (s *service) UpdateCampaign(inputId CampaignDetailInput, inputData CampaignCreateInput) (Campaign, error) {
	campaign, err := s.repository.FindById(inputId.Id)
	if err != nil {
		return campaign, err
	}

	if campaign.UserId != inputData.User.Id {
		return campaign, errors.New("User not owner in campaign")
	}

	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount

	updateCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return updateCampaign, err
	}
	return updateCampaign, nil
}

func (s *service) SaveCampaignImage(input CreateCampaignImageInput, filelocation string) (CampaignImage, error) {
	campaign, err := s.repository.FindById(input.CampaignId)
	if err != nil {
		return CampaignImage{}, err
	}

	if campaign.UserId != input.User.Id {
		return CampaignImage{}, errors.New("User not owner in campaign")
	}

	if input.IsPrimary {
		_, err := s.repository.MarkAllImagesAsNonPrimary(input.CampaignId)
		if err != nil {
			return CampaignImage{}, fmt.Errorf("Failed to create image: %v", err)
		}
	}

	campaignImage := CampaignImage{}
	campaignImage.CampaignId = input.CampaignId

	isPrimary := 0
	if input.IsPrimary {
		isPrimary = 1
	}

	campaignImage.IsPrimary = isPrimary
	campaignImage.FileName = filelocation

	newCampaignImage, err := s.repository.CreateImage(campaignImage)
	if err != nil {
		return newCampaignImage, err
	}
	return newCampaignImage, nil
}
