package middleware

import (
	"fmt"
	"net/http"

	"github.com/Denterry/SocialNetwork/mainService/internal/config"
	"github.com/Denterry/SocialNetwork/mainService/util"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, err := util.ExtractTokenID(ctx, cfg)
		if err != nil {
			fmt.Println(err)
			ctx.String(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}
		ctx.Set("userID", userID)
		ctx.Next()
	}
}
