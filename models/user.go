package models

type User struct {
	Model

	Email  string
	IDCard string
	Verify int
}

func AddUser(data map[string]interface{}) error {
	user := User{
		Email:  data["email"].(string),
		IDCard: data["id_card"].(string),
		Verify: data["verify"].(int),
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}
