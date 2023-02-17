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

func (v *Vote) Check() bool {
	return models.CheckVote(v.ThemeID, v.UserID)
}

func (v *Vote) GetVoteByThemeID() ([]*models.Vote, error) {
	return models.GetVoteByThemeID(v.ThemeID)
}
