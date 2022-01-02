package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) (string, error) {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
		return password, err
	}
	return string(hashpassword), nil
}

func Checkpassword(password, hashpassword string) (string, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password))
	msg := ""
	if err != nil {
		msg = "wrong password"
	}
	return msg, err
}
