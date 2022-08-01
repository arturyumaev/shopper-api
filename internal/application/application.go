package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	users_http "github.com/arturyumaev/shopper-api/internal/domains/user/delivery/http"
	"github.com/arturyumaev/shopper-api/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type IApplication interface {
	Run() error
}

type application struct {
	config *models.Config
}

func (app *application) Run() error {
	initDatabase()

	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	users_http.RegisterHTTPEndpoints(router)

	httpServer := &http.Server{
		Addr:           ":" + app.config.Server.Port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatalf("failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return httpServer.Shutdown(ctx)
}

func initDatabase() {
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")

	dbString := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_HOST, POSTGRES_DB)

	conn, err := pgx.Connect(context.Background(), dbString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
}

func NewApplication(config *models.Config) IApplication {
	return &application{
		config,
	}
}
