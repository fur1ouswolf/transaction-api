package repository

import (
	"github.com/fur1ouswolf/transaction-api/internal/model"
	"time"
)

type TransactionRepository interface {
	CreateTransaction(transaction model.Transaction) error
	GetAllTransactions() ([]model.Transaction, error)
	GetAllTransactionsByTime(startTime, endTime time.Time) ([]model.Transaction, error)
	GetVotes() ([]model.Transaction, error)
}
