package utills

import (
	"errors"
	"net/url"
)

func DecodeUrl(s string) (string, error) {
	decodeString, err := url.QueryUnescape(s)
	if err != nil {
		return "", errors.New("Invalid string")
	}
	return decodeString, nil
}
