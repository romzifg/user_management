package usermodule

import (
	"github.com/gin-gonic/gin"
	"github.com/romzifg/user_management/middleware"
)

func Routes(r *gin.Engine) {
	route := r.Group("/api/user").Use(middleware.ApiTokenMiddleware)

	route.GET("/", ShowAll)
	route.GET("/:id", ShowById)
}