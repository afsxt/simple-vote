package theme_service

import (
	"github.com/afsxt/simple-vote/models"
)

type Theme struct {
	ID          int
	Name        string
	Description string
	State       int
	CreatedBy   string
	ModifiedBy  string
}

func (t *Theme) ExistByName() (bool, error) {
	return models.ExistThemeByName(t.Name)
}

func (t *Theme) ExistByID() (bool, error) {
	return models.ExistThemeByID(t.ID)
}

func (t *Theme) Add() error {
	theme := map[string]interface{}{
		"name":        t.Name,
		"description": t.Description,
	}

	if err := models.AddTheme(theme); err != nil {
		return err
	}

	return nil
}

func (t *Theme) ChangeState() error {
	return models.ChangeThemeState(t.ID, t.State)
}
