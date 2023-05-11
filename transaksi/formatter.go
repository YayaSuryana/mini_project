package transaksi

import "time"

type KampanyeTransaksiFormatter struct {
	ID 				int 		`json:"id"`
	Name 			string 		`json:"name"`
	Amount 			int 		`json:"amount"`
	CreatedAt 		time.Time	`json:"created_at"`
}

// single object (detail transaksi kampanye)
func FormatKampanyeTransaksi(transaksi Transaksi) KampanyeTransaksiFormatter{
	format := KampanyeTransaksiFormatter{}
	format.ID = transaksi.ID
	format.Name = transaksi.User.Name
	format.Amount = transaksi.Amount
	format.CreatedAt = transaksi.CreatedAt
	return format
}

// prural object / jamak
func FormatKampanyeTransaksis(transaksis []Transaksi) []KampanyeTransaksiFormatter{
	if len(transaksis) == 0 {
		return []KampanyeTransaksiFormatter{}
	}

	var transaksisFormatter []KampanyeTransaksiFormatter

	for _, transaksi := range transaksis{
		formatter := FormatKampanyeTransaksi(transaksi)
		transaksisFormatter = append(transaksisFormatter, formatter)
	}
	return transaksisFormatter
}

type UserTransaksiFormatter struct{
	ID			int 	`json:"id"`
	Amount		int 	`json:"amount"`
	Status 		string	`json:"status"`
	CreatedAt	time.Time `json:"created_at"`
	Kampanye 	KampanyeFormatter	`json:"kampanye"`
}

type KampanyeFormatter struct{
	Name		string		`json:"name"`
	ImageURL	string		`json:"image_url"`
}

func FormatterUserTransaksi(transaksi Transaksi) UserTransaksiFormatter {
	formatter := UserTransaksiFormatter{}
	formatter.ID = transaksi.ID
	formatter.Amount = transaksi.Amount
	formatter.Status = transaksi.Status
	formatter.CreatedAt = transaksi.CreatedAt

	kampanyeFormatter := KampanyeFormatter{}
	kampanyeFormatter.Name = transaksi.Kampanye.Name
	kampanyeFormatter.ImageURL = ""
	
	if len(transaksi.Kampanye.KampanyeImages) > 0 {
		kampanyeFormatter.ImageURL = transaksi.Kampanye.KampanyeImages[0].FileName
	}

	formatter.Kampanye = kampanyeFormatter

	return formatter
}
func FormatUserTransaksis(transaksis []Transaksi) []UserTransaksiFormatter{
	if len(transaksis) == 0 {
		return []UserTransaksiFormatter{}
	}

	var transaksisFormatter []UserTransaksiFormatter

	for _, transaksi := range transaksis{
		formatter := FormatterUserTransaksi(transaksi)
		transaksisFormatter = append(transaksisFormatter, formatter)
	}
	return transaksisFormatter
}