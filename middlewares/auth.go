package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserContext struct {
	Username string
}

func CheckAuth(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, ok := claims["id"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("UserID", userId)
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
}
