package auth

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
}
