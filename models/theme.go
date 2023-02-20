package models

import (
	"github.com/jinzhu/gorm"
)

type Theme struct {
	Model
	Name        string
	Description string
	Flag        int
}

// ExistByName 根据名称是否存在
func ExistThemeByName(name string) (bool, error) {
	var t Theme
	err := db.Select("id").Where("name = ? AND deleted_on = ? ", name, 0).First(&t).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if t.ID > 0 {
		return true, nil
	}

	return false, nil
}

// ExistByName id判断是否存在
func ExistThemeByID(id int) (bool, error) {
	var t Theme
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&t).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if t.ID > 0 {
		return true, nil
	}

	return false, nil
}

// AddTheme 增加一个选举主题
func AddTheme(data map[string]interface{}) error {
	t := Theme{
		Name:        data["name"].(string),
		Description: data["description"].(string),
	}

	if err := db.Create(&t).Error; err != nil {
		return err
	}

	return nil
}

// ChangeThemeState 更改选举状态
func ChangeThemeState(id, state int) error {
	return db.Model(&Theme{}).Where("id = ?", id).Update("flag", state).Error
}
