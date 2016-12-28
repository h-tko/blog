package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Blog struct {
	gorm.Model
	Title       string
	Body        string
	IsShow      bool
	ReleaseDate time.Time
}

func (Blog) TableName() string {
	return "blogs"
}

func NewBlog() *Blog {
	return new(Blog)
}

func (blog *Blog) FindById(id int) {
	db.Where("is_show = ?", true).First(&blog, id)
}

func (blog *Blog) FindList(limit int) []*Blog {
	var blogs []*Blog
	db.Where("is_show = ?", true).Order("release_date desc").Limit(limit).Find(&blogs)

	return blogs
}

func RegistBlog(blog Blog) {
	db.NewRecord(blog)
	db.Create(&blog)
	db.Save(&blog)
}
