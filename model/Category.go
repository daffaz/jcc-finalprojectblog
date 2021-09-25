package Model

import "jccblog/Config"

// GetCategory... Fetch all category data
func GetCategory(category *[]Category) error {
	if err := Config.DB.Find(category).Error; err != nil {
		return err
	}
	return nil
}

// CreateCategory... Create a new category
func CreateCategory(category *Category) error {
	if err := Config.DB.Create(category).Error; err != nil {
		return err
	}
	return nil
}
