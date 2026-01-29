package routes

import (
	"miniauth/controllers"
	"miniauth/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *controllers.AuthController)  {
	r.POST("/register", c.Register)
	r.POST("login", c.Login)

	auth := r.Group("/profile")
	auth.Use(middlewares.AuthMiddleware())
	auth.GET("/profile", c.Profile)
}
