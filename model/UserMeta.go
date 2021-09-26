package Model

import (
	"jccblog/Config"
)

// CreateUser... Create a new user
func CreateUserMeta(user *UserMeta) error {
	if err := Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserById... Fetch a single user by it's user_id
func GetUserMetaById(user *UserMeta, id string) error {
	if err := Config.DB.Where("user_id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}
