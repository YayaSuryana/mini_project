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
	Kampanye 	kampanye.Kampanye
	User		user.User
	CreatedAt	time.Time
	UpdatedAt	time.Time
}