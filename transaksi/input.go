package transaksi

import "yayasuryana/user"

type GetKampanyeTrasaksiInput struct{
	ID 			int  		`uri:"id" binding:"required"`
	User		user.User	
}

type CreateTransaksiInput struct{
	KampanyeID		int			`json:"kampanye_id" binding:"required"`
	Amount			int			`json:"amount" binding:"required"`
	User			user.User
}