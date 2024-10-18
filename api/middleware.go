package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/simplebank/token"
)

func authMiddleware(token token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizatioHeader := ctx.GetHeader("authorization")
		if len(authorizatioHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "auhtorization header is not provided"})
			return
		}

		payload, terr := token.VerifyToken(authorizatioHeader)
		if terr != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(terr))
			return
		}

		ctx.Set("authpayload", payload)
		ctx.Next()
	}
}
