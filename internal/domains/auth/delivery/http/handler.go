package http

import (
	"net/http"

	"github.com/arturyumaev/shopper-api/internal/domains/auth"
	"github.com/gin-gonic/gin"
)

type handler struct {
	useCase auth.UseCase
}

func (h *handler) SignUp(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (h *handler) SignIn(c *gin.Context) {
	c.Status(http.StatusOK)
}

func NewHandler(uc auth.UseCase) auth.Handler {
	return &handler{uc}
}
