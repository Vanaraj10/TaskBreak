package routes

import (
	"github.com/Vanaraj10/taskmorph-backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	api := router.Group("/auth")
	{
		api.POST("register", controllers.Register)
		api.POST("login", controllers.Login)
	}

}