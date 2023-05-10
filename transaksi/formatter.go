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