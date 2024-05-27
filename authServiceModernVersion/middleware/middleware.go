package middleware

import (
	"fmt"
	"net/http"

	"github.com/Denterry/SocialNetwork/authServiceModernVersion/internal/config"
	"github.com/Denterry/SocialNetwork/authServiceModernVersion/util"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := util.TokenValid(ctx, cfg)
		if err != nil {
			fmt.Println(err)
			ctx.String(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
