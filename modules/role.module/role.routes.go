package rolemodule

import (
	"github.com/gin-gonic/gin"
	"github.com/romzifg/user_management/middleware"
)

func Routes(r *gin.Engine) {
	route := r.Group("/api/role").Use(middleware.ApiTokenMiddleware)

	route.GET("/", ShowAll)
	route.GET("/:id", ShowById)
	route.POST("/", Create)
	route.PUT("/:id", Update)
	route.DELETE("/:id", Delete)
}