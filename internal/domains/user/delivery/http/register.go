package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func RegisterHTTPEndpoints(router *gin.Engine) {
	userEndpoints := router.Group("/users")
	{
		userEndpoints.POST("/", test)
		userEndpoints.GET("/", test)
		userEndpoints.GET("/:id", test)
		userEndpoints.PUT("/:id", test)
		userEndpoints.DELETE("/:id", test)
	}
}
