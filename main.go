package main

import (
	"github.com/gin-gonic/gin"

	"example/gin_jwt/controllers"
	"example/gin_jwt/middlewares"
	"example/gin_jwt/models"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.0.103"})

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")
}
