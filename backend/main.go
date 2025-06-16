package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Vanaraj10/taskmorph-backend/config"
	"github.com/Vanaraj10/taskmorph-backend/middleware"
	"github.com/Vanaraj10/taskmorph-backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()

	c.AddFunc("@every 10m", func() {
		fmt.Println("Running every 5 minutes at", time.Now())
	})

	c.Start()
	startServer()
}

func startServer () {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	config.ConnectDB()
	fmt.Println("TaskMorph Backend is running...")
	router := gin.Default()

	// Add CORS middleware
	router.Use(middleware.CORSMiddleware())

	routes.SetupRoutes(router)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"message": "TaskMorph Backend is running",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	router.Run("0.0.0.0:" + port)
}
