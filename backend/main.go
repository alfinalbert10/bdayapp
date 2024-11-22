package main

import (
	"bdayappbknd/database"
	"bdayappbknd/handlers"
	"bdayappbknd/managers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

func init() {
	database.Initialize()
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	userManager := managers.NewUserManager()                // Pointer
	userHandler := handlers.NewUserHandlerFrom(userManager) // Use pointer directly
	userHandler.RegisterUserApis(router)

	// listen and serve on 0.0.0.0:8080
	port := os.Getenv("PORT") // Use the PORT environment variable provided by Render

	if port == "" {
		port = "8080" // Default to port 8080 if no environment variable is set
	}

	// Ensure binding to the correct port
	router.Run(":" + port)
}
