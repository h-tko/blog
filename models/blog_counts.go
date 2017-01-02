package models

import (
	"github.com/jinzhu/gorm"
)

type BlogCount struct {
	gorm.Model
	BlogId    uint
	GoodCount int
}

func (BlogCount) TableName() string {
	return "blog_counts"
}

func NewBlogCount() *BlogCount {
	return new(BlogCount)
}

func RegistBlogCount(bc BlogCount) {
	db.NewRecord(bc)
	db.Create(&bc)
	db.Save(&bc)
}
