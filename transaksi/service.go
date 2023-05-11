package transaksi

import (
	"errors"
	"yayasuryana/kampanye"
	"yayasuryana/payment"
)

type Service interface{
	GetTransaksiByKampanyeID(input GetKampanyeTrasaksiInput) ([]Transaksi, error)
	GetTransaksiByUserID(userID int) ([]Transaksi, error)
	CreateTransaksi(input CreateTransaksiInput) (Transaksi, error)
}

type service struct{
	repository Reqpository
	kampanyeRepository kampanye.Repository
	paymentService payment.Service
}

func NewService(repository Reqpository, kampanyeRepository kampanye.Repository, paymentService payment.Service) *service{
	return &service{repository, kampanyeRepository, paymentService}
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

// transaksi by user id
func (s *service) GetTransaksiByUserID(userID int) ([]Transaksi, error){
	transaksi, err := s.repository.GetByUserID(userID)
	if err != nil {
		return transaksi, err
	}
	return transaksi, nil
}

// create new transaksi 
func (s *service) CreateTransaksi(input CreateTransaksiInput) (Transaksi, error){
	transaksi := Transaksi{}
	transaksi.KampanyeID = input.KampanyeID
	transaksi.Amount = input.Amount
	transaksi.UserID = input.User.ID
	transaksi.Status = "pending"


	newTransaksi, err := s.repository.Save(transaksi)
	if err != nil {
		return newTransaksi, err
	}

	// mapping input Transaksi ke entity payment
	paymentTransaksi := payment.Transaksi{
		ID:     newTransaksi.ID,
		Amount: newTransaksi.Amount,
	}
	
	// panggil payment service
	payment, err := s.paymentService.GetPaymentURL(paymentTransaksi, input.User)
	if err != nil {
		return newTransaksi, err
	}

	newTransaksi.PaymentURL = payment

	newTransaksi, err = s.repository.Update(newTransaksi)
	if err != nil {
		return newTransaksi, err
	}

	return newTransaksi, nil
}