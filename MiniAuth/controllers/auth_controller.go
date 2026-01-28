package controllers

import (
	"miniauth/models"
	"miniauth/services"
	"miniauth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)
type AuthController struct {
	Service *services.AuthService
}

// func (c *AuthController) Register(ctx *gin.Context) {
// 	var user models.User
// 	ctx.ShouldBindJSON(&user)

// 	if err := c.Service.Register(&user); err != nil{
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
// 		return
// 	}
// 	ctx.JSON(200, gin.H{"message":"registered succesfully"})
// }

func (c *AuthController) Register(ctx *gin.Context) {
	var req models.RegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := c.Service.Register(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "registered successfully",
	})
}

// func (c AuthController) Login(ctx *gin.Context) {
// 	var req struct{
// 		Email string `json:"email"`
// 		Password string `json:"password"`
// 	}
// 	ctx.ShouldBindJSON(&req)

// 	user, err := c.Service.Login(req.Email, req.Password)
// 	if err != nil {
// 		ctx.JSON(http.StatusUnauthorized, gin.H{"error":err.Error()})
// 		return
// 	}
// 	token, _ := utils.GenerateToken(user.ID)
// 	ctx.JSON(200, gin.H{"token":token,})
// }
func (c *AuthController) Login(ctx *gin.Context) {
	var req models.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.Service.Login(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, _ := utils.GenerateToken(user.ID)
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}


func (c *AuthController) Profile(ctx *gin.Context) {
	userID := ctx.MustGet("user_id")
	ctx.JSON(200, gin.H{"user_id":userID})
}
