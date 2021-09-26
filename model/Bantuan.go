package Model

import (
	"jccblog/Config"
)

// GetBantuan... Fetch all bantuan data
func GetBantuan(bantuan *[]Bantuan) error {
	if err := Config.DB.Find(bantuan).Error; err != nil {
		return err
	}
	return nil
}

// CreateBantuan... Create a new bantuan
func CreateBantuan(bantuan *Bantuan) error {
	if err := Config.DB.Create(bantuan).Error; err != nil {
		return err
	}
	return nil
}
