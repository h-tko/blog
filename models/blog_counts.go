package models

import (
	"time"
)

type BlogCount struct {
	BlogID    uint `gorm:"primary_key"`
	GoodCount int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (BlogCount) TableName() string {
	return "blog_counts"
}

func NewBlogCount() *BlogCount {
	return new(BlogCount)
}

func (model *BlogCount) Regist() {
	db.NewRecord(model)
	db.Create(&model)
}

func (BlogCount) IncrementGood(blogID uint) int {
	blogCount := BlogCount{BlogID: blogID}

	db.First(&blogCount)
	blogCount.GoodCount++
	db.Save(&blogCount)

	return blogCount.GoodCount
}
