package main

import (
	"jccblog/Config"
	"jccblog/Model"
	"jccblog/Routes"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())), &gorm.Config{})

	if err != nil {
		log.Fatal("Status:", err)
	}

	Config.DB.AutoMigrate(&Model.Blog{})
	Config.DB.AutoMigrate(&Model.Category{})
	Config.DB.AutoMigrate(&Model.User{})
	r := Routes.SetupRouter()
	r.Run(":10000")
}
