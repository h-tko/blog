package models

import (
	"github.com/jinzhu/gorm"
)

type Blog struct {
	gorm.Model
	Title  string
	Body   string
	IsShow bool
}

func (Blog) TableName() string {
	return "blogs"
}

func NewBlog() *Blog {
	return new(Blog)
}

func (blog *Blog) FindById(id int) *Blog {
	db.First(&blog, id)

	return blog
}

func (blog *Blog) FindAll() []*Blog {
	var blogs []*Blog
	db.Where("is_show = ?", true).Find(&blogs)

	return blogs
}
