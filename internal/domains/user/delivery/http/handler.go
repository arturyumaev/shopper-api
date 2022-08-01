package http

import (
	"github.com/arturyumaev/shopper-api/internal/domains/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	useCase user.UseCase
}

func (h *handler) CreateUser(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (h *handler) GetUser(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (h *handler) GetUsers(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (h *handler) UpdateUser(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (h *handler) DeleteUser(c *gin.Context) {
	c.Status(http.StatusOK)
}

func NewHandler(useCase user.UseCase) user.Handler {
	return &handler{useCase}
}
