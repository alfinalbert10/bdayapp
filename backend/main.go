package main

import (
	"bdayappbknd/database"
	"bdayappbknd/handlers"
	"bdayappbknd/managers"
	"github.com/gin-gonic/gin"
	"os"
	
)

func init() {
	database.Initialize()
}

func main() {
	router := gin.Default()

	userManager := managers.NewUserManager()                // Pointer
	userHandler := handlers.NewUserHandlerFrom(userManager) // Use pointer directly
	userHandler.RegisterUserApis(router)
	// router.Run() // listen and serve on 0.0.0.0:8080
	port := os.Getenv("PORT") // Use the PORT environment variable

	if port == "" {
		port = "8080" // Default port if none is set
	}

	// Run the server on the specified port
	router.Run(":" + port)
}
