package handler

import (
	"net/http"
	"yayasuryana/helper"
	"yayasuryana/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct{
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler{
	return &userHandler{userService}
}

// Handler Register
func (h *userHandler) RegisterUser(c *gin.Context){
	// Tangkap Inputan User
	// Map inputan dari user ke struct Register User Input
	// struct di atas kita passing sebagai parameter ke service
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	// pengecekan validate
	if err != nil {
		errors := helper.FormatValidationError(err)
		// mapping dari setiap errors
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Gagal membuat akun", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil{
		response := helper.APIResponse("Gagal membuat akun", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest,response)
		return
	}

	fomatter := user.FormatUser(newUser, "tokentokentoken")
	response := helper.APIResponse("Berhasil membuat akun", http.StatusOK, "success", fomatter)

	c.JSON(http.StatusOK, response)
}

// handler login
func (h *userHandler) Login(c *gin.Context){
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMesaage := gin.H{"errors": errors}
		response := helper.APIResponse("Login Gagal", http.StatusUnprocessableEntity, "error", errorMesaage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loginUser, err := h.userService.Login(input)

	if err != nil{
		errorMesaage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Email atau password salah", http.StatusBadRequest, "error", errorMesaage)
		
		c.JSON(http.StatusBadRequest,response)
		return
	}
	formatter := user.FormatUser(loginUser, "token token token")
	response := helper.APIResponse("Berhasil Login", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

// handler chek email
func (h *userHandler) CheckEmail(c *gin.Context){
	var input user.CheckEmailAvailable
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMesaage := gin.H{"errors": errors}
		resoponse := helper.APIResponse("Cek email gagal", http.StatusUnprocessableEntity, "error", errorMesaage)
		c.JSON(http.StatusUnprocessableEntity, resoponse)
		return
	}

	checkEmail, err := h.userService.CheckEmail(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.APIResponse("Cek Email Gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": checkEmail,
	}

	// membuat metaMessage di response
	var metaMessage string
	if checkEmail {
		metaMessage = "Email is available"
	}else{
		metaMessage = "Email not available"
	}


	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}
