package main

import (
	"log"
	"miniauth/config"
	"miniauth/controllers"
	"miniauth/models"
	"miniauth/repositories"
	"miniauth/routes"
	"miniauth/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

)


func main() {
	// laod env 
	if err := godotenv.Load();  err != nil{
		log.Fatal("error loading .env failed")
	}

	db := config.ConnectDB()
	db.AutoMigrate(&models.User{})

	userRepo := &repositories.UserRepository{DB: db}
	authService := &services.AuthService{UserRepo: userRepo}
	authController := &controllers.AuthController{Service: authService}

	r := gin.Default()
	routes.RegisterRoutes(r, authController)
	r.Run(":8000")
}