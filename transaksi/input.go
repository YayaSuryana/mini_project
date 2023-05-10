package transaksi

import "yayasuryana/user"

type GetKampanyeTrasaksiInput struct{
	ID 			int  		`uri:"id" binding:"required"`
	User		user.User	
}