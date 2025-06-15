package routes

import (
	"github.com/Vanaraj10/taskmorph-backend/controllers"
	"github.com/Vanaraj10/taskmorph-backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Auth routes
	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	// AI routes
	ai := router.Group("/ai")
	{
		ai.POST("/breakdown", controllers.BreakdownTask)
	}

	// Protected task routes
	tasks := router.Group("/tasks")
	tasks.Use(middleware.AuthMiddleware())
	{
		tasks.POST("/create", controllers.CreateTask)
		tasks.GET("/", controllers.GetTasks)
		tasks.GET("/:id", controllers.GetTask)
		tasks.PATCH("/:taskID/step/:stepID/complete", controllers.CompleteStep)
		tasks.DELETE("/:id", controllers.DeleteTask)
	}
}
