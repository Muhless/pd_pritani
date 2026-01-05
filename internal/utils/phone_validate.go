package utils

import (
	"errors"
	"regexp"
)

func ValidatePhone(phone string) error {
	re := regexp.MustCompile(`^(?:\+62|0)\d{8,13}$`) // +62xxxxxxxx or 0xxxxxxx
	if !re.MatchString(phone) {
		return errors.New("phone number is not valid")
	}
	return nil
}
