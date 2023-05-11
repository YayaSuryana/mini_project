package transaksi

import (
	"time"
	"yayasuryana/kampanye"
	"yayasuryana/user"
)

type Transaksi struct {
	ID 			int
	KampanyeID 	int
	UserID		int
	Amount 		int
	Status		string
	Code 		string
	PaymentURL 	string
	User       	user.User
	Kampanye   	kampanye.Kampanye
	CreatedAt	time.Time
	UpdatedAt	time.Time
}