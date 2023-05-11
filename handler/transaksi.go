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

// Get User Transaksi
// handler
// ambil nilai user dari jwt/middleware
// service
// repository => ambil data transaksi (preload kampanye)

func (h *transaksiHandler) GetUserTransaksi(c *gin.Context){
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	userTransaksi, err := h.service.GetTransaksiByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Gagal menampilkan user transaksi", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return	
	}

	response := helper.APIResponse("User transaksi", http.StatusOK, "success", transaksi.FormatUserTransaksis(userTransaksi))
	c.JSON(http.StatusOK, response)
}