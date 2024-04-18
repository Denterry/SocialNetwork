package main

import (
	"github.com/Denterry/SocialNetwork/authService/internal/handlers"
	"github.com/Denterry/SocialNetwork/authService/internal/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	storage.InitDb()
	defer storage.Db.Close()

	router := gin.Default()

	router.POST("/register", handlers.RegisterUser)
	router.PUT("/update/:username", handlers.UpdateUser)
	router.POST("/login", handlers.LoginUser)

	router.POST("/post/create", handlers.CreatePost)
	router.PUT("/post/update", handlers.UpdatePost)
	router.DELETE("/post/delete", handlers.DeletePost)
	router.GET("/post/get/:post_id", handlers.GetPost)
	router.GET("/post/asdasdasdasdasdasdasdasdasdasdasdasdasd", handlers.GetListPosts)

	router.Run(":8080")
}
