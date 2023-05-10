package transaksi

import (
	"errors"
	"yayasuryana/kampanye"
)

type Service interface{
	GetTransaksiByKampanyeID(input GetKampanyeTrasaksiInput) ([]Transaksi, error)
}

type service struct{
	repository Reqpository
	kampanyeRepository kampanye.Repository
}

func NewService(repository Reqpository, kampanyeRepository kampanye.Repository) *service{
	return &service{repository, kampanyeRepository}
}
// transaksi by kampanye id
func (s *service) GetTransaksiByKampanyeID(input GetKampanyeTrasaksiInput) ([]Transaksi, error){
	// get kampanye
	// check kampanye.userID != user_id_yang_melakukan_request
	kampanye, err := s.kampanyeRepository.FindByID(input.ID)
	if err != nil {
		return []Transaksi{}, err
	}

	if kampanye.UserID != input.User.ID{
		return []Transaksi{}, errors.New("Akses ditolak kamu bukan owner dari kampanye ini")
	}
	transaksi, err := s.repository.GetByKampanyeID(input.ID)
	if err != nil {
		return transaksi, err
	}
	return transaksi, nil
}