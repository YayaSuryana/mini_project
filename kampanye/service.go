package kampanye

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetKampanye(userID int) ([]Kampanye, error)
	GetKampanyeByID(input GetKampanyeDetailInput) (Kampanye, error)
	CreateKampanye(input CreateKampanyeInput) (Kampanye, error)
	UpdateKampanye(inputID GetKampanyeDetailInput, inputData CreateKampanyeInput) (Kampanye, error)
	SaveKampanyeImage(input CreateKampanyeImage, fileLocation string) (KampanyeImage, error)
}

type service struct {
	repository Repository
}

// // UpdateKampanye implements Service
// func (*service) UpdateKampanye(inputID GetKampanyeDetailInput, inputData CreateKampanyeInput) (Kampanye, error) {
// 	panic("unimplemented")
// }

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetKampanye(userID int) ([]Kampanye, error) {
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

func (s *service) GetKampanyeByID(input GetKampanyeDetailInput) (Kampanye, error) {
	kampanye, err := s.repository.FindByID(input.ID)
	if err != nil {
		return kampanye, err
	}
	return kampanye, nil
}

// create kampanye
func (s *service) CreateKampanye(input CreateKampanyeInput) (Kampanye, error) {
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
	if err != nil {
		return newKampanye, err
	}

	return newKampanye, nil
}

func (s *service) UpdateKampanye(inputID GetKampanyeDetailInput, inputData CreateKampanyeInput) (Kampanye, error) {
	// tangkap id dari uri
	kampanye, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return kampanye, err
	}

	// cek apakah userID sama dengan id user yang melakukan requrest?
	if kampanye.UserID != inputData.User.ID{
		return kampanye, errors.New("Akses ditolak, kamu bukan owner dari kampanye ini!")
	}
	// mapping ke struct input
	kampanye.Name = inputData.Name
	kampanye.ShortDescription = inputData.ShortDescription
	kampanye.Description = inputData.Description
	kampanye.GoalAmount = inputData.GoalAmount
	kampanye.Perks = inputData.Perks

	updateKampanye, err := s.repository.Update(kampanye)
	if err != nil {
		return updateKampanye, err
	}

	return updateKampanye, nil

}

// save kampanye image
func (s *service) SaveKampanyeImage(input CreateKampanyeImage, fileLocation string) (KampanyeImage, error){
	kampanye, err := s.repository.FindByID(input.KampanyeID)
	if err != nil {
		return KampanyeImage{}, err
	}
	if kampanye.UserID != input.User.ID {
		return KampanyeImage{}, errors.New("Akses ditolak, kamu bukan owner dari kampanye ini!")
	}
	isPrimary := 0
	if input.IsPrimary {
		isPrimary = 1
		_, err := s.repository.MarkNonIsPrimary(input.KampanyeID)
		if err != nil {
			return	KampanyeImage{}, err
		}
	}

	kampanyeImage := KampanyeImage{}

	kampanyeImage.KampanyeID = input.KampanyeID
	kampanyeImage.IsPrimary = isPrimary
	kampanyeImage.FileName = fileLocation

	newKampanyeImage, err := s.repository.CreateImage(kampanyeImage)
	if err != nil {
		return newKampanyeImage, err
	}

	return newKampanyeImage, nil
}
