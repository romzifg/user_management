package rolemodule

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/romzifg/user_management/models"
	"gorm.io/gorm"
)

func ShowAll(c *gin.Context) {
	var roles []models.Role

	models.DB.Find(&roles)
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message": "Success",
		"data": roles,
	})
}

func ShowById(c *gin.Context) {
	var role models.Role
	id := c.Param("id")

	if err := models.DB.First(&role, id).Error; err != nil {
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
		"data": role,
	})
}

func Create(c *gin.Context) {
	var role models.Role

	if err := c.ShouldBindJSON(&role); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"message": "Bad Request",
			"data": nil,
		})
	}

	models.DB.Create(&role)
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message": "Success",
		"data": role,
	})
}

func Update(c *gin.Context) {
	var role models.Role
	var dto UpdateRoleDto
	id := c.Param("id")
	
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"message": "Bad Request",
			"data": nil,
		})
	}

	if err := models.DB.First(&role, id).Error; err != nil {
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

	models.DB.Model(&role).Where("id = ?", id).Updates(&dto)	
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message": "Success",
		"data": role,
	})
}

func Delete(c *gin.Context) {
	var role models.Role
	id := c.Param("id")

	if err := models.DB.First(&role, id).Error; err != nil {
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

	models.DB.Delete(&role).Where("id = ?", id)
	c.JSON(http.StatusNotFound, gin.H{
		"statusCode": http.StatusOK,
		"message": "Success",
		"data": id,
	})
}