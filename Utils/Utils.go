package Utils

import (
	"errors"
	"jccblog/Config"
	"jccblog/Model"
	"net/url"
)

var Admin []Model.User
var Writer []Model.User

func ValidateUrl(link string) (string, error) {
	_, err := url.ParseRequestURI(link)

	if err != nil {
		return "", errors.New("link tidak valid")
	}

	return link, nil
}

func GetAdminAccount() {
	Config.DB.Raw("SELECT * FROM user WHERE role = ?", "admin").Scan(&Admin)
}

func GetWriterAccount() {
	Config.DB.Raw("SELECT * FROM user WHERE role = ?", "writer").Scan(&Writer)
}
