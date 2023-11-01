package campaign

type Service interface {
	FindCampaign(userId int) ([]Campaign, error)
	FindCampaignById(input CampaignDetailInput) (Campaign, error)
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
