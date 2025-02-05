package middlewares

import (
	"abramed_go/cmd/api/helpers"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

func LogAcesso(c *gin.Context) {

	ip := c.GetHeader("x-forwarded-for")
	if ip == "" {
		ip = c.GetHeader("x-real-ip")
	}
	fmt.Println(time.Now(), "RemoteIP: ", c.RemoteIP(), "IP: ", ip, "UserAgent: ", c.Request.UserAgent())
	c.Next()
}

func AuthenticationMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {

		fmt.Println("user")
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

func CorsMiddleware(c *gin.Context) {

	allowedOrigin := os.Getenv("ALLOWED_ORIGIN")
	if allowedOrigin == "" {
		allowedOrigin = "http://localhost:3000" // Valor padrão
	}

	// Verifica a origem da requisição
	origin := c.Request.Header.Get("Origin")
	equal, err := regexp.MatchString(allowedOrigin, origin)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !equal {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "Origin not allowed",
		})
		return
	}

	// Métodos permitidos
	allowedMethods := os.Getenv("ALLOWED_METHODS")
	if allowedMethods == "" {
		allowedMethods = "GET, POST, PUT, DELETE, PATCH, OPTIONS" // Métodos padrão
	}

	// Cabeçalhos permitidos
	allowedHeaders := os.Getenv("ALLOWED_HEADERS")
	if allowedHeaders == "" {
		allowedHeaders = "Content-Type, Authorization" // Cabeçalhos padrão
	}

	// Define os cabeçalhos CORS
	c.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
	c.Writer.Header().Set("Access-Control-Allow-Methods", allowedMethods)
	c.Writer.Header().Set("Access-Control-Allow-Headers", allowedHeaders)

	// Permite requisições OPTIONS (preflight)
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.Next()
}
