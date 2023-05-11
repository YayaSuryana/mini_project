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
// repo mencari data transaksi sutu kampanye

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

// input nominal dari user 
// handler tangkap input terus di mapping ke struct input
// panggil service buat transaksi, manggil sistem midtrans
// panggil repository create new transaksi data
func (h *transaksiHandler) CreateTransaksi(c *gin.Context){
	var input transaksi.CreateTransaksiInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Gagal membuat transaksi baru-baru", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newTransaksi, err := h.service.CreateTransaksi(input)

	if err != nil {
		response := helper.APIResponse("Gagal membuat transaksi baru baru ini", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	response := helper.APIResponse("Berhasil membuat transaksi baru", http.StatusOK, "success", transaksi.FormatTransaksi(newTransaksi))
	c.JSON(http.StatusOK, response)
}