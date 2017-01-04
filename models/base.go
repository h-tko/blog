package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pelletier/go-toml"
)

var db *gorm.DB

func InitDatabase() error {
	config, err := toml.LoadFile("./config/app.toml")

	if err != nil {
		return err
	}

	host := config.Get("database.host").(string)
	user := config.Get("database.user").(string)
	password := config.Get("database.password").(string)
	port := config.Get("database.port").(string)
	dbname := config.Get("database.dbname").(string)
	dbsource, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password))

	if err != nil {
		return err
	}

	db = dbsource
	db.LogMode(true)

	return nil
}

func CloseDatabase() {
	db.Close()
}
