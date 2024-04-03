package main

import (
	"auth/internal/handlers"
	"auth/internal/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	storage.InitDb()
	defer storage.Db.Close()

	router := gin.Default()

	router.POST("/register", handlers.RegisterUser)
	router.PUT("/update/:username", handlers.UpdateUser)
	router.POST("/login", handlers.LoginUser)

	router.Run(":8080")
}
