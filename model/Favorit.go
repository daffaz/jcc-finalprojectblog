package Model

import (
	"jccblog/Config"
)

func CreateFavorit(favorit *Favorit) error {
	if err := Config.DB.Create(favorit).Error; err != nil {
		return err
	}
	return nil
}

func GetFavoritByUID(favorit *[]Favorit, id string) error {
	if err := Config.DB.Where("user_id = ?", id).Find(favorit).Error; err != nil {
		return err
	}
	return nil
}

func DeleteFavorit(favorit *Favorit, id string) error {
	Config.DB.Where("id = ?", id).Delete(favorit)
	return nil
}
