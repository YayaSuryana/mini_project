package kampanye

import (
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface{
	GetKampanye(userID int) ([]Kampanye, error)
	GetKampanyeByID(input GetKampanyeDetailInput) (Kampanye, error)
	CreateKampanye(input CreateKampanyeInput) (Kampanye, error)
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

// create kampanye
func (s *service) CreateKampanye(input CreateKampanyeInput) (Kampanye, error){
	kampanye := Kampanye{}
	kampanye.Name = input.Name
	kampanye.ShortDescription = input.ShortDescription
	kampanye.Description = input.Description
	kampanye.GoalAmount = input.GoalAmount
	kampanye.Perks = input.Perks
	kampanye.UserID = input.User.ID

	// membuat slug dengan package slug (https://github.com/gosimple/slug)
	slugString := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	kampanye.Slug = slug.Make(slugString)
	newKampanye, err := s.repository.Save(kampanye)
	if err != nil{
		return newKampanye, err
	}

	return newKampanye, nil
}