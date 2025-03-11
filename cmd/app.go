package cmd

import (
	"log"
	"os"

	"github.com/JerryJeager/Symptomify-Backend/manualwire"
	"github.com/JerryJeager/Symptomify-Backend/middleware"
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
	tabController := manualwire.GetTabController()

	api := router.Group("/api/v1")
	users := api.Group("/users")
	tabs := api.Group("tabs")

	users.POST("/signup", userController.CreateUser)
	users.POST("/verify", userController.VerifyUser)
	users.POST("/login", userController.Login)
	users.GET("", middleware.JwtAuthMiddleware(), userController.GetUser)

	tabs.POST("", middleware.JwtAuthMiddleware(), tabController.CreateTab)
	tabs.GET("", middleware.JwtAuthMiddleware(), tabController.GetTabs)
	tabs.DELETE("/:tab_id", middleware.JwtAuthMiddleware(), tabController.DeleteTab)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
