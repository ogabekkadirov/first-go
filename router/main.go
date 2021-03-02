package router

import (
	"github.com/gin-gonic/gin"
)


func Init() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("api/v1")
	{
		v1.GET("/test", func(c *gin.Context) {
			return
		})
	}

	return router
}