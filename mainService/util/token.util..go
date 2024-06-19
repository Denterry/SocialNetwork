package util

import (
	"fmt"
	"strings"
	"time"

	"github.com/Denterry/SocialNetwork/mainService/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func GenerateToken(user_id uuid.UUID, cfg *config.Config) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id.String()
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(cfg.JWT.TTL)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(cfg.JWT.Secret))
}

func TokenValid(ctx *gin.Context, cfg *config.Config) error {
	tokenString := ExtractToken(ctx)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(cfg.JWT.Secret), nil
	})
	if err != nil {
		return err
	}

	return nil
}

func ExtractToken(ctx *gin.Context) string {
	token := ctx.Query("token")
	if token != "" {
		return token
	}

	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(ctx *gin.Context, cfg *config.Config) (uuid.UUID, error) {
	tokenString := ExtractToken(ctx)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.JWT.Secret), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userIDStr, ok := claims["user_id"].(string) // Extract as string
		if !ok {
			return uuid.Nil, fmt.Errorf("user_id is not a string")
		}

		uid, err := uuid.Parse(userIDStr) // Parse string to UUID
		if err != nil {
			return uuid.Nil, fmt.Errorf("user_id is not a valid UUID: %v", err)
		}

		return uid, nil
	}

	return uuid.Nil, fmt.Errorf("invalid token or claims")
}
