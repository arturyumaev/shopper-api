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

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	r.Run(fmt.Sprintf(":%s", app.config.Server.Port))
}

func NewApplication(config *models.Config) IApplication {
	return &application{
		config,
	}
}
