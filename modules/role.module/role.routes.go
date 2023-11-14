package rolemodule

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	route := r.Group("/api/role")

	route.GET("/", ShowAll)
	route.GET("/:id", ShowById)
	route.POST("/", Create)
	route.PUT("/:id", Update)
	route.DELETE("/:id", Delete)
}