package Model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Categories string `json:"categories"`
}

func (b *Category) TableName() string {
	return "category"
}
