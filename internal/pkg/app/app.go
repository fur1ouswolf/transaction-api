package app

import (
	"github.com/fur1ouswolf/transaction-api/internal/repository"
	"github.com/fur1ouswolf/transaction-api/internal/transaction/handler"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"os"
)

type ginApp struct {
	logger *slog.Logger
	router *gin.Engine
}

func NewGinApp(repository repository.TransactionRepository, logger *slog.Logger) App {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	api := router.Group(os.Getenv("API_PREFIX"))

	handler.NewGinTransactionHandler(api, logger, repository)

	return &ginApp{
		logger: logger,
		router: router,
	}
}

func (g *ginApp) Start() error {
	g.logger.Info("Starting server...")
	return g.router.Run(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT"))
}

func (g *ginApp) Stop() error {
	// TODO: add graceful shutdown
	return nil
}
