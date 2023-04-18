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