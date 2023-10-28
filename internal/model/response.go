package model

type ResultResponse struct {
	CandidateID int `json:"candidate_id"`
	VoteCount   int `json:"vote_count"`
}
