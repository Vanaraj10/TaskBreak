package main

import (
	"fmt"
	"os"

	"github.com/Vanaraj10/taskmorph-backend/config"
	"github.com/Vanaraj10/taskmorph-backend/middleware"
	"github.com/Vanaraj10/taskmorph-backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	fmt.Printf("🌟 Server running on all interfaces at port %s\n", port)
	fmt.Printf("🌐 Local access: http://localhost:%s\n", port)
	fmt.Printf("🌐 Network access: http://192.168.94.124:%s\n", port)

	router.Run("0.0.0.0:" + port)
}
