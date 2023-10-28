package handler

import (
	"github.com/fur1ouswolf/transaction-api/internal/model"
	"github.com/fur1ouswolf/transaction-api/internal/repository"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"time"
)

type transactionHandler struct {
	logger     *slog.Logger
	repository repository.TransactionRepository
}

func NewGinTransactionHandler(api *gin.RouterGroup, logger *slog.Logger, repository repository.TransactionRepository) {
	handler := &transactionHandler{
		logger:     logger.WithGroup("TransactionHandler"),
		repository: repository,
	}
	api.POST("/transaction/ballot", handler.AddBallotTransaction)
	api.POST("/transaction/vote", handler.AddVoteTransaction)
	api.POST("/transaction/result", handler.AddResultTransaction)
	api.GET("/transactions", handler.GetTransactions)
	api.GET("/transactions/time", handler.GetTransactionsByTime)
	api.GET("/result", handler.GetResult)
}

// AddBallotTransaction godoc
// @Summary Add ballot transaction
// @Tags transaction
// @Description Add ballot transaction
// @Accept json
// @Produce json
// @Param ballot_id body int true "ballot_id"
// @Success 200 {string} string "success"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 400 {string} string "ballot_id is required"
// @Failure 500 {string} string "Internal server error"
// @Router /transaction/ballot [post]
func (h *transactionHandler) AddBallotTransaction(ctx *gin.Context) {
	var transaction model.BallotTransaction
	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	if transaction.BallotID == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ballot_id is required"})
		return
	}
	transaction.Type = "ballot"
	transaction.CreatedAt = time.Now()

	if err := h.repository.CreateTransaction(&transaction); err != nil {
		h.logger.Error("Failed to create transaction", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// AddVoteTransaction godoc
// @Summary Add vote transaction
// @Tags transaction
// @Description Add vote transaction
// @Accept json
// @Produce json
// @Param ballot_id body int true "ballot_id"
// @Param candidate_id body int true "candidate_id"
// @Success 200 {string} string "success"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 400 {string} string "ballot_id is required"
// @Failure 400 {string} string "candidate_id is required"
// @Failure 500 {string} string "Internal server error"
// @Router /transaction/vote [post]
func (h *transactionHandler) AddVoteTransaction(ctx *gin.Context) {
	var transaction model.VoteTransaction
	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	if transaction.BallotID == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ballot_id is required"})
		return
	}
	if transaction.CandidateID == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "candidate_id is required"})
		return
	}
	transaction.Type = "vote"
	transaction.CreatedAt = time.Now()

	if err := h.repository.CreateTransaction(&transaction); err != nil {
		h.logger.Error("Failed to create transaction", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// AddResultTransaction godoc
// @Summary Add result transaction
// @Tags transaction
// @Description Add result transaction
// @Accept json
// @Produce json
// @Param candidate_id body int true "candidate_id"
// @Param vote_count body int true "vote_count"
// @Success 200 {string} string "success"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 400 {string} string "candidate_id is required"
// @Failure 400 {string} string "vote_count is required"
// @Failure 500 {string} string "Internal server error"
// @Router /transaction/result [post]
func (h *transactionHandler) AddResultTransaction(ctx *gin.Context) {
	var transaction model.ResultTransaction
	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	if transaction.CandidateID == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "candidate_id is required"})
		return
	}
	if transaction.VoteCount == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "vote_count is required"})
		return
	}

	transaction.Type = "result"
	transaction.CreatedAt = time.Now()

	if err := h.repository.CreateTransaction(&transaction); err != nil {
		h.logger.Error("Failed to create transaction", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// GetTransactions godoc
// @Summary Get all transactions
// @Tags transaction
// @Description Get all transactions
// @Accept json
// @Produce json
// @Success 200 {object} []model.Transaction
// @Failure 500 {string} string "Internal server error"
// @Router /transactions [get]
func (h *transactionHandler) GetTransactions(ctx *gin.Context) {
	transactions, err := h.repository.GetAllTransactions()
	if err != nil {
		h.logger.Error("Failed to get transactions", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

// GetTransactionsByTime godoc
// @Summary Get all transactions between startTime and endTime
// @Tags transaction
// @Description Get all transactions between startTime and endTime
// @Accept json
// @Produce json
// @Param start_time body string true "start_time"
// @Param end_time body string true "end_time"
// @Success 200 {object} []model.Transaction
// @Failure 400 {string} string "Invalid JSON"
// @Failure 400 {string} string "Invalid start_time"
// @Failure 400 {string} string "Invalid end_time"
// @Failure 500 {string} string "Internal server error"
// @Router /transactions/time [get]
func (h *transactionHandler) GetTransactionsByTime(ctx *gin.Context) {
	var fields map[string]string
	err := ctx.ShouldBindJSON(&fields)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	startTime, err := time.Parse("2006-01-02T15:04:05.999MST", fields["start_time"])
	if err != nil {
		h.logger.Error("Failed to bind JSON", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_time"})
		return
	}
	endTime, err := time.Parse("2006-01-02T15:04:05.999MST", fields["end_time"])
	if err != nil {
		h.logger.Error("Failed to bind JSON", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_time"})
		return
	}

	transactions, err := h.repository.GetAllTransactionsByTime(startTime, endTime)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

// GetResult godoc
// @Summary Get result
// @Tags transaction
// @Description Get result
// @Accept json
// @Produce json
// @Success 200 {object} []model.ResultResponse
// @Failure 500 {string} string "Internal server error"
// @Router /result [get]
func (h *transactionHandler) GetResult(ctx *gin.Context) {
	voteTransactions, err := h.repository.GetVotes()
	if err != nil {
		h.logger.Error("Failed to get votes", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	resultCnt := make(map[int]int)
	for _, t := range voteTransactions {
		resultCnt[*t.(*model.VoteTransaction).CandidateID]++
	}
	result := make([]model.ResultResponse, 0, len(resultCnt))
	for key, value := range resultCnt {
		result = append(result, model.ResultResponse{CandidateID: key, VoteCount: value})
	}

	ctx.JSON(http.StatusOK, gin.H{"result": result})
}
