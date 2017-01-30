package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type CategoryCount struct {
	Category      uint
	Name          string
	CategoryCount int
}

type Blog struct {
	gorm.Model
	Title       string
	Body        string
	Keywords    string
	IsShow      bool
	Category    uint
	ReleaseDate time.Time
	BlogCount   BlogCount `gorm:"ForeignKey:BlogID"`
}

func (Blog) TableName() string {
	return "blogs"
}

func NewBlog() *Blog {
	return new(Blog)
}

func (blog *Blog) FindByID(id int) {
	db.Where("is_show = ?", true).Where("release_date <= now()").First(&blog, id)

	db.Model(&blog).Related(&blog.BlogCount, "BlogCount")
}

func (blog *Blog) FindPopularList(limit int) []*Blog {
	var blogs []*Blog
	db.Where("is_show = ?", true).Where("release_date <= now()").Joins("inner join blog_counts on blogs.id = blog_counts.blog_id").Order("blog_counts.good_count desc").Limit(limit).Find(&blogs)

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

func (blog *Blog) FindListByCategory(category uint) []*Blog {
	var blogs []*Blog
	db.Where("is_show = ?", true).Where("release_date <= now()").Where("category = ?", category).Order("release_date desc").Find(&blogs)

	return blogs
}

func (blog *Blog) FindListOrderCategory() []*Blog {
	var blogs []*Blog
	db.Where("is_show = ?", true).Where("release_date <= now()").Order("category asc, release_date desc").Find(&blogs)

	return blogs
}

func (blog *Blog) CategoryCount() []*CategoryCount {
	var categoryCounts []*CategoryCount
	db.Table("blogs").Select("category, categories.name as name, count(*) as category_count").Group("category, categories.name").Where("is_show = ?", true).Where("release_date <= now()").Joins("inner join categories on blogs.category = categories.id").Scan(&categoryCounts)

	return categoryCounts
}

func (model *Blog) Regist() {
	db.NewRecord(model)
	db.Create(&model)
}

func (model *Blog) Update() {
	db.Save(&model)
}
