package transaksi

import (
	"time"
	"yayasuryana/user"
)

type Transaksi struct {
	ID 			int
	KampanyeID 	int
	UserID		int
	Amount 		int
	Status		string
	Code 		string
	User		user.User
	CreatedAt	time.Time
	UpdatedAt	time.Time
}