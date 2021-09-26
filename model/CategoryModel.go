package Model

// Category represents the model for a category
type Category struct {
	ID           int
	CategoryName string `json:"category_name"`
}

func (b *Category) TableName() string {
	return "category"
}
