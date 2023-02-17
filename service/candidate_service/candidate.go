package candidate_service

import (
	"github.com/afsxt/simple-vote/models"
)

type Candidates struct {
	ID          int
	Name        string
	Description string
	ThemeID     int
	CreatedBy   string
	ModifiedBy  string
}

func (c *Candidates) ExistBy() (bool, error) {
	return models.ExistCandidateBy(c.Name, c.ThemeID)
}

func (c *Candidates) Add() error {
	candidate := map[string]interface{}{
		"name":        c.Name,
		"description": c.Description,
		"theme_id":    c.ThemeID,
	}

	if err := models.AddCandidate(candidate); err != nil {
		return err
	}

	return nil
}

func (c *Candidates) GetCount() (int, error) {
	return models.GetCandidateCountByThemeID(c.ThemeID)
}
