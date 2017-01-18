package models

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name string
	Sort uint
}

func (Category) TableName() string {
	return "categories"
}

func NewCategory() *Category {
	return new(Category)
}

func (category *Category) All() []*Category {
	var categories []*Category
	db.Order("sort asc").Find(&categories)

	return categories
}
