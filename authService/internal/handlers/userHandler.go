package handlers

import (
	"auth/internal/storage"
	"auth/pkg/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register a new user
func RegisterUser(g *gin.Context) {
	fmt.Println("dasdasdasdas")
	var regUser models.LogSignIn
	if err := g.ShouldBindJSON(&regUser); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(regUser.Username, regUser.Password)
	_, err := storage.Db.Exec("INSERT INTO account(username, password) VALUES ($1, $2)", regUser.Username, regUser.Password)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Update an existing user
func UpdateUser(g *gin.Context) {
	userUsername := g.Param("username")

	var user models.User
	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := storage.Db.Exec("UPDATE account SET name=$1, surname=$2, birthday=$3, email=$4, phone=$5 WHERE username=$6",
		user.Name, user.Surname, user.Birthday, user.Email, user.Phone, userUsername)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// Log in a user via username and password
func LoginUser(g *gin.Context) {
	var logUser models.LogSignIn
	if err := g.BindJSON(&logUser); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var password string
	err := storage.Db.QueryRow("SELECT password FROM account WHERE username=$1", logUser.Username).Scan(&password)
	if err != nil {
		g.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if password != logUser.Password {
		g.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "User authenticated successfully"})

	// var logUser models.LogSignIn
	// if err := g.ShouldBindJSON(&logUser); err != nil {
	// 	g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// row := storage.Db.QueryRow("SELECT id FROM users WHERE username=$1 AND password=$2", logUser.Username, logUser.Password)
	// var id int
	// if err := row.Scan(&id); err != nil {
	// 	g.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	// 	return
	// }

	// g.JSON(http.StatusOK, gin.H{"message": "Login successful", "user_id": id})
}
