package main

import (
	"log"
	"net/http"
	"strings"
	"yayasuryana/auth"
	"yayasuryana/handler"
	"yayasuryana/helper"
	"yayasuryana/kampanye"
	"yayasuryana/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main(){
	dsn := "root:@tcp(127.0.0.1:3306)/mini_project?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
			log.Fatal(err.Error())
	}

	 userRepository := user.NewRepository(db)
	 kampanyeRepository := kampanye.NewRepository(db)

	 userService 	:= user.NewService(userRepository)
	 kampanyeService := kampanye.NewService(kampanyeRepository)

	 
	 authService 	:= auth.NewService()
	 userHandler 	:= handler.NewUserHandler(userService, authService)
	 kampanyeHandler := handler.NewKampanyeHandler(kampanyeService)


	 router := gin.Default()
	 api 	:= router.Group("/api/v1")
	 router.Static("/images","./img")
	 api.POST("/users", userHandler.RegisterUser)
	 api.POST("/login", userHandler.Login)
	 api.POST("/email_checkers", userHandler.CheckEmail)
	 api.POST("/avatar", AuthMiddleware(authService, userService), userHandler.UploadAvatar)
	 api.GET("/kampanye", kampanyeHandler.GetKampanye)

	 router.Run()	   
}

func AuthMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func (c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer"){
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return 
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return 
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil{
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return 
		}
		c.Set("currentUser", user)
	}	
}