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
	BlogCount   BlogCount `gorm:"ForeignKey:BlogID"`
}

func (Blog) TableName() string {
	return "blogs"
}

func NewBlog() *Blog {
	return new(Blog)
}

func (blog *Blog) FindById(id int) {
	db.Where("is_show = ?", true).Where("release_date <= now()").First(&blog, id)

	db.Model(&blog).Related(&blog.BlogCount, "BlogCount")
}

func (blog *Blog) FindPopularList(limit int) []*Blog {
	var blogs []*Blog
	db.Where("is_show = ?", true).Where("release_date <= now()").Joins("inner join blog_counts on blogs.id = blog_counts.blog_id").Order("blog_counts.good_count").Limit(limit).Find(&blogs)

	for _, data := range blogs {
		db.Model(&data).Related(&data.BlogCount, "BlogCount")
	}

	return blogs
}

func (blog *Blog) FindList(limit int) []*Blog {
	var blogs []*Blog
	db.Where("is_show = ?", true).Where("release_date <= now()").Order("release_date desc").Limit(limit).Find(&blogs)

	return blogs
}

func RegistBlog(blog *Blog) {
	db.NewRecord(blog)
	db.Create(&blog)
	db.Save(&blog)
}
