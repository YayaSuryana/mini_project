package kampanye

type Service interface{
	GetKampanye(userID int) ([]Kampanye, error)
	GetKampanyeByID(input GetKampanyeDetailInput) (Kampanye, error)
}

type service struct{
	repository Repository
}

func NewService(repository Repository) *service{
	return &service{repository}
}

func (s *service) GetKampanye(userID int) ([]Kampanye, error){
	if userID != 0 {
		user, err := s.repository.FindByUserID(userID)
		if err != nil {
			return user, err
		}
		return user, nil
	}
	
	user, err := s.repository.FindAll()
	if err != nil {
		return user, err
	}

	return user, nil
}

func(s *service) GetKampanyeByID(input GetKampanyeDetailInput) (Kampanye, error) {
	kampanye, err := s.repository.FindByID(input.ID)
	if err != nil{
		return kampanye, err
	}
	return kampanye, nil
}