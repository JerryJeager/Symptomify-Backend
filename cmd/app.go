package cmd

import (
	"log"
	"os"

	"github.com/JerryJeager/Symptomify-Backend/middleware"
	"github.com/JerryJeager/Symptomify-Backend/manualwire"
	"github.com/gin-gonic/gin"
)

func ExecuteApiRoutes() {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to Symptomify",
		})
	})

	userController := manualwire.GetUserController()

	api := router.Group("/api/v1")
	users := api.Group("/users")

	users.POST("/signup", userController.CreateUser)
	users.POST("/verify", userController.VerifyUser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
