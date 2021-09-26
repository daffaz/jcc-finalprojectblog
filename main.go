package main

import (
	"jccblog/Config"
	"jccblog/Model"
	"jccblog/Routes"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

// @title Final Project | Blog API with Gin, Gorm and Gosimple/Slug
// @version 1.0.0
// @description This Project is still production ready and still receive support from developer. For admin account (username : daffaz, password : admin), writer account (username : william, password: shakespear)
// @termsOfService http://swagger.io/terms/

// @contact.name Daffa Zaky
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	Config.DB, err = gorm.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())), &gorm.Config{})

	if err != nil {
		log.Fatal("Status:", err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	Config.DB.AutoMigrate(&Model.Blog{})
	Config.DB.AutoMigrate(&Model.Category{})
	Config.DB.AutoMigrate(&Model.User{})
	Config.DB.AutoMigrate(&Model.Comment{})
	Config.DB.AutoMigrate(&Model.Bantuan{})
	Config.DB.AutoMigrate(&Model.UserMeta{})
	Config.DB.AutoMigrate(&Model.Favorit{})
	r := Routes.SetupRouter()
	r.Run(":" + port)
}
