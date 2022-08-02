package http

import (
	"net/http"

	"github.com/arturyumaev/shopper-api/internal/domains/user"
	"github.com/arturyumaev/shopper-api/models"
	"github.com/gin-gonic/gin"
	logrus "github.com/sirupsen/logrus"
)

type handler struct {
	useCase user.UseCase
}

func (h *handler) CreateUser(c *gin.Context) {
	user := &models.User{}
	if err := c.BindJSON(user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.useCase.CreateUser(c.Request.Context(), user); err != nil {
		logrus.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}

func (h *handler) GetUser(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (h *handler) GetUsers(c *gin.Context) {
	users, err := h.useCase.GetUsers(c.Request.Context())
	if err != nil {
		logrus.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
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
