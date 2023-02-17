package models

import "github.com/jinzhu/gorm"

type Candidates struct {
	Model
	Name        string //同一个主题下候选人name不能重复
	Description string
	ThemeID     int
}

// ExistCandidateBy 判断主题候选人是否存在
func ExistCandidateBy(name string, theme_id int) (bool, error) {
	var c Candidates
	err := db.Select("id").Where("name = ? AND theme_id = ? AND deleted_on = ? ", name, theme_id, 0).First(&c).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if c.ID > 0 {
		return true, nil
	}

	return false, nil
}

// AddCandidate　增加一个候选人
func AddCandidate(data map[string]interface{}) error {
	candidates := Candidates{
		Name:        data["name"].(string),
		Description: data["description"].(string),
		ThemeID:     data["theme_id"].(int),
	}

	if err := db.Create(&candidates).Error; err != nil {
		return err
	}

	return nil
}

func GetCandidateCountByThemeID(themeID int) (int, error) {
	var count int
	if err := db.Model(&Candidates{}).Where("theme_id = ?", themeID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
