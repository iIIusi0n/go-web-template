package middlewares

import (
	"log"

	"api-server/controllers/auth"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		log.Println("No token found: ", err)

		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	claims, err := auth.ParseToken(token)
	if err != nil {
		log.Println("Invalid token: ", err)

		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	c.Set("uid", claims.UserID)
	c.Set("name", claims.Name)
	c.Next()
}
