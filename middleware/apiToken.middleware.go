package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ApiTokenMiddleware(c *gin.Context) {
	requiredToken := os.Getenv("REQUIRED_TOKEN")
	token := c.Request.Header.Get("api_token")

	if token == "" || token != requiredToken {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"statusCode": http.StatusUnauthorized,
			"message": "Unauthorized",
		})
	}

	c.Next()
}