package application

import (
	"context"
	"fmt"
	"github.com/arturyumaev/shopper-api/internal/domains/user"
	userHttp "github.com/arturyumaev/shopper-api/internal/domains/user/delivery/http"
	userRepo "github.com/arturyumaev/shopper-api/internal/domains/user/repository/postgres"
	userUC "github.com/arturyumaev/shopper-api/internal/domains/user/usecase"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/arturyumaev/shopper-api/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type IApplication interface {
	Run() error
}

type application struct {
	httpServer *http.Server
	config     *models.Config

	userUC user.UseCase
}

func (app *application) Run() error {
	go func() {
		if err := app.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return app.httpServer.Shutdown(ctx)
}

func initPostgresDatabase() *pgx.Conn {
	var (
		POSTGRES_USER     = os.Getenv("POSTGRES_USER")
		POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
		POSTGRES_HOST     = os.Getenv("POSTGRES_HOST")
		POSTGRES_DB       = os.Getenv("POSTGRES_DB")
	)

	dbString := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_HOST, POSTGRES_DB)

	conn, err := pgx.Connect(context.Background(), dbString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Successfully connected to database: %v\n", dbString)
	defer conn.Close(context.Background())

	return conn
}

func NewApplication(config *models.Config) IApplication {
	conn := initPostgresDatabase()

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

	userRepository := userRepo.NewRepository(conn)
	userUseCase := userUC.NewUseCase(userRepository)
	userHttp.RegisterHTTPEndpoints(router, userUseCase)

	httpServer := &http.Server{
		Addr:           ":" + config.Server.Port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &application{
		httpServer,
		config,
		userUseCase,
	}
}
