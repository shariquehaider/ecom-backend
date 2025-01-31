package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/shariquehaider/ecom-backend/utils"
)

func VerifyTokenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			ctx.Abort()
			return
		}
		onlyToken := strings.TrimPrefix(tokenString, "Bearer ")
		token, err := jwt.Parse(onlyToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, gin.Error{
					Type: gin.ErrorTypePublic,
					Meta: "Invalid signing method",
				}
			}
			return utils.JwtSecret, nil
		})

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx.Set("_id", claims["_id"])
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
		}
	}
}
