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
func(h *kampanyeHandler) GetKampanyes(c *gin.Context){
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

func (h *kampanyeHandler) GetKampanye(c *gin.Context){
	// handler : mapping id yang di url ke struct input => service, call formatter
	// service : inputnya struct input => menangkap id di url, memanggil repo
	// repository : get kampanye by id

	var input kampanye.GetKampanyeDetailInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.APIResponse("Gagal memuat detail kampanye", http.StatusBadRequest, "error",nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	kampanyeDetail, err := h.service.GetKampanyeByID(input)
	if err != nil{
		response := helper.APIResponse("Gagal memuat detail kampanye", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Kampanye detail", http.StatusOK, "success", kampanye.FormatKampanyeDetail(kampanyeDetail))
	c.JSON(http.StatusOK,response)
}

