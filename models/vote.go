package models

type Vote struct {
	Model

	ThemeID     int
	UserID      int
	CandidateID int
}

func AddVote(data map[string]interface{}) error {
	v := Vote{
		ThemeID:     data["theme_id"].(int),
		UserID:      data["user_id"].(int),
		CandidateID: data["candidate_id"].(int),
	}

	if err := db.Create(&v).Error; err != nil {
		return err
	}

	return nil
}

func CheckVote(themeID, userID int) bool {
	return db.Model(&Vote{}).Where("theme_id = ? and user_id = ?", themeID, userID).RecordNotFound()
}

func GetVoteByThemeID(themeID int) ([]*Vote, error) {
	var votes []*Vote
	if err := db.Model(&Vote{}).Where("theme_id = ?", themeID).Find(&votes).Error; err != nil {
		return nil, err
	}
	return votes, nil
}
