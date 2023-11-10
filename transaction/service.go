package transaction

type Service interface {
	GetTransactionByCampaignId(input CampaignTransactionInput) ([]Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTransactionByCampaignId(input CampaignTransactionInput) ([]Transaction, error) {
	transaction, err := s.repository.GetCampaignById(input.Id)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
