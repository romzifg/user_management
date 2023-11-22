package usermodule

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/romzifg/user_management/helpers"
	"github.com/romzifg/user_management/models"
	"gorm.io/gorm"
)

func ShowAll(c *gin.Context) {
	var users []models.User

	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message": "Success",
		"data": users,
	})
}

func ShowById(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := models.DB.First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"statusCode": http.StatusNotFound,
				"message": "Not Found",
				"data": nil,
			})
		default:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"statusCode": http.StatusBadRequest,
				"message": "Bad Request",
				"data": nil,
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message": "Success",
		"data": user,
	})
}

func Create(c *gin.Context) {
	var user CreateUserDataDto

	err :=  c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	// Hash Password
	hash, _ := helpers.HashPassword(user.Password)

	// Save to table user
	createUser := models.User{FirstName: user.FirstName, LastName: user.LastName, Address: user.Address, Phone: user.Phone}
	models.DB.Create(&createUser)

	// Save to table auth
	createAuth := models.Auth{Username: user.Username, Email: user.Email, Password: hash, RoleId: user.RoleId}
	models.DB.Create(&createAuth)

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message": "Success",
		"data": user,
	})
}