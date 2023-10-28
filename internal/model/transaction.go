package model

import (
	repoModel "github.com/fur1ouswolf/transaction-api/internal/repository/transaction/model"
	"time"
)

type Transaction interface {
	ToRepo() *repoModel.RepoTransaction
}

type BaseTransaction struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`
	Signature string    `json:"signature"`
	CreatedAt time.Time `json:"created_at"`
}

type BallotTransaction struct {
	BaseTransaction
	BallotID *int `json:"ballot_id"`
}

type VoteTransaction struct {
	BaseTransaction
	BallotID    *int `json:"ballot_id"`
	CandidateID *int `json:"candidate_id"`
}

type ResultTransaction struct {
	BaseTransaction
	CandidateID *int `json:"candidate_id"`
	VoteCount   *int `json:"vote_count"`
}

func (t *BallotTransaction) ToRepo() *repoModel.RepoTransaction {
	return &repoModel.RepoTransaction{
		ID:        t.ID,
		Type:      t.Type,
		CreatedAt: t.CreatedAt,
		BallotID:  t.BallotID,
	}
}

func (t *VoteTransaction) ToRepo() *repoModel.RepoTransaction {
	return &repoModel.RepoTransaction{
		ID:          t.ID,
		Type:        t.Type,
		BallotID:    t.BallotID,
		CandidateID: t.CandidateID,
	}
}

func (t *ResultTransaction) ToRepo() *repoModel.RepoTransaction {
	return &repoModel.RepoTransaction{
		ID:          t.ID,
		Type:        t.Type,
		CandidateID: t.CandidateID,
		VoteCount:   t.VoteCount,
	}
}
