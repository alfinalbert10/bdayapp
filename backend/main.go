package main

import (
	"bdayappbknd/database"
	"bdayappbknd/handlers"
	"bdayappbknd/managers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	// router.Run() // listen and serve on 0.0.0.0:8080
	//port := os.Getenv("PORT") // Use the PORT environment variable
	//
	//if port == "" {
	//	port = "8080" // Default port if none is set
	//}

	// Run the server on the specified port
	router.Run()
}
