package http

import (
	"github.com/arturyumaev/shopper-api/internal/domains/auth"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc auth.UseCase) {
	h := NewHandler(uc)

	userEndpoints := router.Group("/auth")

	{
		userEndpoints.POST("/sign-up", h.SignUp)
		userEndpoints.POST("/sign-in", h.SignIn)
	}
}
