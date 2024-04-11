package jwt

import (
	"time"

	"github.com/Denterry/SocialNetwork/postService/internal/domain/models"
	"github.com/golang-jwt/jwt/v5"
)

func NewToken(post models.Post, app models.App, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authid"] = post.AuthorID
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["appid"] = app.ID

	tokenString, err := token.SignedString([]byte(app.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
