package application

import (
	"fmt"
	"net/http"

	"github.com/arturyumaev/shopper-api/models"
	"github.com/gin-gonic/gin"
)

type IApplication interface {
	Run()
}

type application struct {
	config *models.Config
}

func (app *application) Run() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	r.Run(fmt.Sprintf("%s:%s", app.config.Server.Host, app.config.Server.Port))
}

func NewApplication(config *models.Config) IApplication {
	return &application{
		config,
	}
}
