package Utils

import (
	"errors"
	"net/url"
)

func ValidateUrl(link string) (string, error) {
	_, err := url.ParseRequestURI(link)

	if err != nil {
		return "", errors.New("link tidak valid")
	}

	return link, nil
}
