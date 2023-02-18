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

func CheckVote(themeID, userID int) (bool, error) {
	var vote Vote
	r := db.Model(&Vote{}).Where("theme_id = ? and user_id = ?", themeID, userID).Find(&vote)
	return r.RowsAffected > 0, r.Error
}

type VoteResult struct {
	CandidateID int
	VoteCount   int
}

func GetVotes(data map[string]interface{}) ([]*VoteResult, error) {
	var result []*VoteResult
	v := Vote{
		ThemeID:     data["theme_id"].(int),
		UserID:      data["user_id"].(int),
		CandidateID: data["candidate_id"].(int),
	}
	conn := db.Model(&Vote{})
	if v.ThemeID != 0 {
		conn = conn.Where("theme_id = ?", v.ThemeID)
	}
	if v.CandidateID != 0 {
		conn = conn.Where("candidate_id = ?", v.CandidateID)
	}
	if err := conn.Select("candidate_id, count(user_id) as vote_count").Group("candidate_id").Scan(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func GetVoteUsers(pageNum int, pageSize int, data map[string]interface{}) ([]*User, error) {
	var result []*User
	v := Vote{
		ThemeID:     data["theme_id"].(int),
		UserID:      data["user_id"].(int),
		CandidateID: data["candidate_id"].(int),
	}
	conn := db.Model(&Vote{}).Offset(pageNum).Limit(pageSize)
	if v.ThemeID != 0 {
		conn = conn.Where("theme_id = ?", v.ThemeID)
	}
	if v.CandidateID != 0 {
		conn = conn.Where("candidate_id = ?", v.CandidateID)
	}
	if err := conn.Select("vote_user.*").Joins("left join vote_user on user_id = vote_user.id").Scan(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
