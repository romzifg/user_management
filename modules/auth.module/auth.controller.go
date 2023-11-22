package authmodule

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/romzifg/user_management/helpers"
	"github.com/romzifg/user_management/models"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	var auth models.Auth
	var authDto AuthDto

	err :=  c.ShouldBindJSON(&authDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	err = models.DB.First(&auth, "email = ?", authDto.Email).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"statusCode": http.StatusNotFound,
				"message": "Not Found",
				"data": nil,
			})
			return
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"statusCode": http.StatusBadRequest,
				"message": "Bad Request",
				"data": nil,
			})
			return
		}
	}

	_, errCompare := helpers.ComparePassword(auth.Password, authDto.Password)
	if errCompare != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
				"statusCode": http.StatusUnauthorized,
				"message": "Email or Password is not match",
				"data": nil,
			})
			return
	}

	token, err := helpers.GenerateToken(int(auth.Id))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"statusCode": http.StatusUnauthorized,
			"message": "Email or Password is not match",
			"data": nil,
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 120, "", "localhost:4001", false, true)

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message": "Success",
		"token": token,
	})
}