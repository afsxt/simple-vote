package user_service

import (
	"regexp"

	"github.com/afsxt/simple-vote/models"
	"github.com/afsxt/simple-vote/pkg/util"
)

type User struct {
	ID         int
	Email      string
	IDCard     string
	Verify     int
	CreatedBy  string
	ModifiedBy string
}

func (u *User) CheckValid() (bool, error) {
	if ok := util.IsEmailValid(u.Email); !ok {
		return false, nil
	}
	match, err := regexp.MatchString(`\b[A-Z]\d{6}[\(][(0-9][\)]`, u.IDCard)
	if err != nil {
		return false, err
	}
	return match, nil
}

func (u *User) Add() error {
	user := map[string]interface{}{
		"email":   u.Email,
		"id_card": u.IDCard,
		"verify":  u.Verify,
	}

	if err := models.AddUser(user); err != nil {
		return err
	}

	return nil
}
