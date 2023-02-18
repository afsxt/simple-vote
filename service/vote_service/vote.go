package vote_service

import (
	"github.com/afsxt/simple-vote/models"
)

type Vote struct {
	ID          int
	ThemeID     int
	UserID      int
	CandidateID int
	CreatedBy   string
	ModifiedBy  string

	PageSize int
	PageNum  int
}

func (v *Vote) Add() error {
	vote := map[string]interface{}{
		"theme_id":     v.ThemeID,
		"user_id":      v.UserID,
		"candidate_id": v.CandidateID,
	}

	if err := models.AddVote(vote); err != nil {
		return err
	}

	return nil
}

func (v *Vote) Check() (bool, error) {
	return models.CheckVote(v.ThemeID, v.UserID)
}

func (v *Vote) GetVote() ([]*models.VoteResult, error) {
	vote := map[string]interface{}{
		"theme_id":     v.ThemeID,
		"user_id":      v.UserID,
		"candidate_id": v.CandidateID,
	}
	return models.GetVotes(vote)
}

func (v *Vote) GetVoteUsers() ([]*models.User, error) {
	vote := map[string]interface{}{
		"theme_id":     v.ThemeID,
		"user_id":      v.UserID,
		"candidate_id": v.CandidateID,
	}
	return models.GetVoteUsers(v.PageNum, v.PageSize, vote)
}
