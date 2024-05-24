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
	// postClient := client.NewPostServiceClient("localhost:8081")
	// postHandler := handlers.NewPostServiceClientAPI(postClient)

	router.POST("/sign-up", handlers.RegisterUser)
	router.POST("/sign-in", handlers.LoginUser)
	router.PUT("/update/:username", handlers.UpdateUser)

	// router.POST("/posts/create", postHandler.CreatePost)
	// router.PUT("/posts/:post_id", postHandler.UpdatePost)
	// router.DELETE("/posts/:post_id", postHandler.DeletePost)
	// router.GET("/posts/get/:post_id", postHandler.GetPost)
	// router.GET("/posts", postHandler.GetListPosts)

	router.Run("localhost:8080")
}
