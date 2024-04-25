package middleware

import (
	"net/http"

	"github.com/sahil-cloud/backend/constants"
	"github.com/sahil-cloud/backend/helper"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientToken := ctx.Request.Header.Get("token")
		if clientToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No Authorization Header Provided"})
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		ctx.Set(constants.USER_NAME, claims.Name)
		ctx.Set(constants.USER_ID, claims.ID)
		ctx.Set(constants.USER_CRYPTO_KEY, claims.CryptoKey)
		ctx.Next()
	}
}
