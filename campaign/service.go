package campaign

type Service interface {
	FindCampaigns(userID int) ([]Campaign, error)
}

type service struct {
	respository Repository
}

func NewService(respository Repository) *service {
	return &service{respository}
}

func (s *service) FindCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.respository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.respository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}
