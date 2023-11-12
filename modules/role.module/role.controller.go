package rolemodule

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/romzifg/user_management/models"
	"gorm.io/gorm"
)

func ShowAll(c *gin.Context) {
	var roles []models.RoleModel

	models.DB.Find(&roles)
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message": "Success",
		"data": roles,
	})
}

func ShowById(c *gin.Context) {
	var role models.RoleModel
	id := c.Param("id")

	if err := models.DB.First(&role, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"statusCode": http.StatusNotFound,
				"message": "Not Found",
				"data": nil,
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message": "Success",
		"data": role,
	})
}