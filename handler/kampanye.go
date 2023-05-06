package handler

import (
	"net/http"
	"strconv"
	"yayasuryana/helper"
	"yayasuryana/kampanye"

	"github.com/gin-gonic/gin"
)

// tangkap parameter di handler (user_id atau none) jika terdapat parameter maka tampilkan kampanye sesuai dengan user_idnya jika none maka tampilkan semua (GET)
// handler ke service
// service menentukan repository mana yang di panggil,
// repository --> FindAll dan FindByUserID
// db

type kampanyeHandler struct{
	service kampanye.Service
}

func NewKampanyeHandler(service kampanye.Service) *kampanyeHandler{
	return &kampanyeHandler{service}
}

//api/v1/kampanye
func(h *kampanyeHandler) GetKampanye(c *gin.Context){
	userID, _ := strconv.Atoi(c.Query("user_id"))

	kampanyes, err := h.service.GetKampanye(userID)
	if err != nil {
		response := helper.APIResponse("Error Get Kampanye", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List Kampanye", http.StatusOK, "success", kampanye.FormatKampanyes(kampanyes))
	c.JSON(http.StatusOK, response)
}

