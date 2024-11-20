package main

import (
	"bdayappbknd/database"
	"bdayappbknd/handlers"
	"bdayappbknd/managers"
	"github.com/gin-gonic/gin"
)

func init() {
	database.Initialize()
}

func main() {
	router := gin.Default()

	userManager := managers.NewUserManager()                // Pointer
	userHandler := handlers.NewUserHandlerFrom(userManager) // Use pointer directly
	userHandler.RegisterUserApis(router)
	router.Run() // listen and serve on 0.0.0.0:8080
}
