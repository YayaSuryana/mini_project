package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"yayasuryana/helper"
	"yayasuryana/kampanye"
	"yayasuryana/user"

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

// tankap parameter dari user input ke struct input
// note : current user diambil dari jwt/handler
// panggil user service parameternya input struct (sekaligus buat slug otomatis)
// panggil repository untuk simpan data kampanye baru

func (h *kampanyeHandler) CreateKampanye(c *gin.Context){
	var input kampanye.CreateKampanyeInput

	err := c.ShouldBindJSON(&input)
	if err != nil{
		errors := helper.FormatValidationError(err)
		errorMesaage := gin.H{"errors": errors}
		resoponse := helper.APIResponse("Data baru gagal dibuat", http.StatusUnprocessableEntity, "error", errorMesaage)
		c.JSON(http.StatusUnprocessableEntity, resoponse)
		return
	}

	// mendapatkan set context (middleware) yang didapatkan dari balikan func authmiddleware dengan bentuk integer
	// key nya adalah currentUser lalu merubah currentUser menjadi user.User
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newKampanye, err := h.service.CreateKampanye(input)

	if err != nil {
		resoponse := helper.APIResponse("Data baru gagal dibuat", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, resoponse)
		return
	}

	resoponse := helper.APIResponse("Data baru berhasil dibuat", http.StatusOK, "success", kampanye.FormatKampanye(newKampanye))
	c.JSON(http.StatusOK, resoponse)
}

// user memasukan inputan
// handler 
// mapping inputan ke inputan struct (inputan update & uri)
// inputan user dan uri di passing ke service
// service (find kampanye by id, tangkap parameter sama ketika input)
// repository update data kampanye

func (h *kampanyeHandler) UpdateKampanye(c *gin.Context){
	var inputID kampanye.GetKampanyeDetailInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Update data kampanye gagal", http.StatusBadRequest, "error",nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}


	// input data
	var inputData kampanye.CreateKampanyeInput
	
	err = c.ShouldBindJSON(&inputData)
	if err != nil{
		errors := helper.FormatValidationError(err)
		errorMesaage := gin.H{"errors": errors}

		response := helper.APIResponse("Update data kampanye gagal", http.StatusUnprocessableEntity, "error", errorMesaage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// set id yang melakukan request agar dilakukan pengecekan di service
	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	updateKampanye, err := h.service.UpdateKampanye(inputID, inputData)

	if err != nil {
		resoponse := helper.APIResponse("Update data kampanye gagal", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, resoponse)
		return
	}

	resoponse := helper.APIResponse("Data kampanye berhasil diupdate", http.StatusOK, "success", kampanye.FormatKampanye(updateKampanye))
	c.JSON(http.StatusOK, resoponse)
}

// handler 
// (tangkap inputan user ubah ke struct input)
// save image ke folder tertentu 
// service (kondisi manggil point 2 di repo, panggil point 1 di repo)
// repository : 
// 1. create image/save ke table kampanye_images
// 2. ubah is_primary true ke false

func (h *kampanyeHandler) UploadImage(c *gin.Context){
	var input kampanye.CreateKampanyeImage

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMesaage := gin.H{"errors": errors}

		response := helper.APIResponse("Upload kampanye image gagal", http.StatusUnprocessableEntity, "error", errorMesaage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	userID := currentUser.ID

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded":false }
		response := helper.APIResponse("Upload kampanye image gagal", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// mendapatkan set context (middleware) yang didapatkan dari balikan func authmiddleware dengan bentuk integer
	// key nya adalah currentUser lalu merubah currentUser menjadi user.User

	path := fmt.Sprintf("img/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	
	if err != nil {
		data := gin.H{"is_uploaded":false}
		response := helper.APIResponse("upload kampanye image gagal", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.SaveKampanyeImage(input, path)
	if  err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("upload kampanye image gagal", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("upload kampanye image berhasil", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}