package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/romzifg/user_management/models"
)

func RequiredAuth(c *gin.Context) {
	// get cookie
	stringToken, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"statusCode": http.StatusUnauthorized,
			"message": "Unauthorized",
		})
	}

	token, _ := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
	
		return []byte(os.Getenv("JWT_TOKEN")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"statusCode": http.StatusUnauthorized,
				"message": "Unauthorized",
			})	
		}

		var auth models.Auth
		models.DB.First(&auth, claims["userId"])

		if auth.Id == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"statusCode": http.StatusUnauthorized,
				"message": "Unauthorized",
			})	
			return
		}

		c.Set("auth", auth)
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"statusCode": http.StatusUnauthorized,
			"message": "Unauthorized",
		})
	}
	
}