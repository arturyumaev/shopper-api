package http

import (
	"github.com/arturyumaev/shopper-api/internal/domains/user"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc user.UseCase) {
	h := NewHandler(uc)

	userEndpoints := router.Group("/users")

	{
		userEndpoints.POST("/", h.CreateUser)
		userEndpoints.GET("/", h.GetUsers)
		userEndpoints.GET("/:id", h.GetUser)
		userEndpoints.PUT("/:id", h.UpdateUser)
		userEndpoints.DELETE("/:id", h.DeleteUser)
	}
}
