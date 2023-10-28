package model

import "time"

type RepoTransaction struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Type        string    `json:"type"`
	Signature   string    `json:"signature"`
	CreatedAt   time.Time `json:"created_at"`
	BallotID    *int      `json:"ballot_id"`
	CandidateID *int      `json:"candidate_id"`
	VoteCount   *int      `json:"vote_count"`
}
