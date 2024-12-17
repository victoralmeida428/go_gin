package middlewares

import (
	"abramed_go/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthenticationMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		c.Abort()
		return
	}

	user, err := helpers.VerifyToken(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.Set("user", user)
	c.Next()

}
