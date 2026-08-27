package middleware

import (
	"erosync/internal/shared/security"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	authHeader := c.GetHeader("authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
		c.AbortWithStatusJSON(401, gin.H{"message": "unauthorized"})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == "" {
		c.AbortWithStatusJSON(401, gin.H{"message": "unauthorized"})
	}

	userId, role, err := security.VerifyToken(tokenString)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"message": "unauthorized"})
	}

	c.Set("userId", userId)
	c.Set("role", role)
	c.Next()
}
