package handler

import (
	"net/http"
	"yayasuryana/helper"
	"yayasuryana/transaksi"
	"yayasuryana/user"

	"github.com/gin-gonic/gin"
)

// parameter uri
// tangkap param mapping ke input struct
// panggil service, input struct sebagai parameter
// service dengan kampanye_id bisa panggil repo
// repo mencari data transaction sutu kampanye

type transaksiHandler struct{
	service transaksi.Service
}

func NewTransaksiHandler(service transaksi.Service) *transaksiHandler{
	return &transaksiHandler{service}
}

// detail transaksi berdasarkan kampanye
func (h *transaksiHandler) GetKampanyeTransaksi(c *gin.Context){
	
	var input transaksi.GetKampanyeTrasaksiInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Gagal menampilkan kampanye transaksi", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transaksis, err := h.service.GetTransaksiByKampanyeID(input)
	if err != nil {
		response := helper.APIResponse("Gagal menampilkan kampanye transaksi", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Kampanye detail transaksi", http.StatusOK, "success", transaksi.FormatKampanyeTransaksis(transaksis))
	c.JSON(http.StatusOK, response)
}