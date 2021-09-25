package Model

import (
	"jccblog/Config"
)

// GetUsers... Fetch all user data
func GetUsers(user *[]User) error {
	if err := Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

// CreateUser... Create a new user
func CreateUser(user *User) error {
	if err := Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserById... Fetch a single user by it's user_id
func GetUserById(user *User, id string) error {
	if err := Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUser... Update a single user by it's id
func UpdateUser(user *User, id string) error {
	Config.DB.Save(user)
	return nil
}

// DeleteUser... Delete a single user by it's id
func DeleteUser(user *User, id string) error {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}
