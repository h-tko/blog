package models

import (
	"github.com/jinzhu/gorm"
)

type AccessLog struct {
	gorm.Model
	IpAddress string
	UserAgent string
}

func (AccessLog) TableName() string {
	return "access_logs"
}

func NewAccessLog() *AccessLog {
	return new(AccessLog)
}

func (model *AccessLog) Regist() {
	db.NewRecord(model)
	db.Create(&model)
}
