package transaction

import (
	"github.com/fur1ouswolf/transaction-api/internal/model"
	repoModel "github.com/fur1ouswolf/transaction-api/internal/repository/transaction/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository() (*Repository, error) {
	// Connect to DB
	db, err := gorm.Open(
		postgres.Open(os.Getenv("DB_CONNECTION_STRING")),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}
	// Migrate models
	if err := db.AutoMigrate(&repoModel.RepoTransaction{}); err != nil {
		return nil, err
	}
	return &Repository{DB: db}, nil
}

func ToModel(t *repoModel.RepoTransaction) model.Transaction {
	switch t.Type {
	case "ballot":
		return &model.BallotTransaction{
			BaseTransaction: model.BaseTransaction{
				ID:        t.ID,
				Type:      t.Type,
				CreatedAt: t.CreatedAt,
			},
			BallotID: t.BallotID,
		}
	case "vote":
		return &model.VoteTransaction{
			BaseTransaction: model.BaseTransaction{
				ID:        t.ID,
				Type:      t.Type,
				CreatedAt: t.CreatedAt,
			},
			BallotID:    t.BallotID,
			CandidateID: t.CandidateID,
		}
	case "result":
		return &model.ResultTransaction{
			BaseTransaction: model.BaseTransaction{
				ID:        t.ID,
				Type:      t.Type,
				CreatedAt: t.CreatedAt,
			},
			CandidateID: t.CandidateID,
			VoteCount:   t.VoteCount,
		}
	default:
		return nil
	}
}

// CreateTransaction creates a new transaction in the database
func (r *Repository) CreateTransaction(transaction model.Transaction) error {
	rt := transaction.ToRepo()
	return r.DB.Create(rt).Error
}

// GetAllTransactions returns all transactions from the database
func (r *Repository) GetAllTransactions() ([]model.Transaction, error) {
	var rt []repoModel.RepoTransaction
	transactions := make([]model.Transaction, 0, 50)
	if err := r.DB.Find(&rt).Error; err != nil {
		return nil, err
	}
	for _, t := range rt {
		transactions = append(transactions, ToModel(&t))
	}

	return transactions, nil
}

// GetAllTransactionsByTime returns all transactions from the database between startTime and endTime
func (r *Repository) GetAllTransactionsByTime(startTime, endTime time.Time) ([]model.Transaction, error) {
	var rt []repoModel.RepoTransaction
	transactions := make([]model.Transaction, 0, 50)
	if err := r.DB.Find(&rt, "created_at BETWEEN ? AND ?", startTime, endTime).Error; err != nil {
		return nil, err
	}
	for _, t := range rt {
		transactions = append(transactions, ToModel(&t))
	}
	return transactions, nil
}

// GetVotes returns all vote transactions from the database
func (r *Repository) GetVotes() ([]model.Transaction, error) {
	var rt []repoModel.RepoTransaction
	var transactions []model.Transaction
	if err := r.DB.Find(&rt, "type = ?", "vote").Error; err != nil {
		return nil, err
	}
	for _, t := range rt {
		transactions = append(transactions, ToModel(&t))
	}
	return transactions, nil
}
