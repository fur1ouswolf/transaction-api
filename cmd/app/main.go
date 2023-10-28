package main

import (
	"github.com/fur1ouswolf/transaction-api/internal/pkg/app"
	transactionRepository "github.com/fur1ouswolf/transaction-api/internal/repository/transaction"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

func main() {
	if err := godotenv.Load("/.env"); err != nil {
		panic(err)
	}

	logger := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	repo, err := transactionRepository.NewRepository()
	if err != nil {
		logger.Error(err.Error())
	}

	ginApp := app.NewGinApp(repo, logger)

	if err := ginApp.Start(); err != nil {
		logger.Error(err.Error())
		return
	}
}
