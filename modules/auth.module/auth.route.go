package authmodule

import (
	"github.com/gin-gonic/gin"
	"github.com/romzifg/user_management/middleware"
)

func Routes(r *gin.Engine) {
	route := r.Group("/api/auth").Use(middleware.ApiTokenMiddleware)

	route.POST("/login", Login)
}