package utils

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password, id string) string {
	pass := []byte(fmt.Sprintf("%s:%s", password, id))

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func CheckPassword(password, id, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(fmt.Sprintf("%s:%s", password, id)))
	return err == nil
}

func CheckStrengthPassword(pass string) bool {
	// (?=.*?[A-Z]) find at least have 1 Upper Case
	// (?=.*?[a-z]) find at least have 1 Lower Case
	// (?=.*?[!@#\$&*~]) find at least have 1 Symbol
	// (?=.*?[0-9]) find at least have 1 number
	// [A-Za-z0-9]{6,} make sure the letter have 8 characters
	secure := true
	tests := []string{".{8,}", "[a-z]", "[A-Z]", "[0-9]"}
	for _, test := range tests {
		t, _ := regexp.MatchString(test, pass)
		if !t {
			secure = false
			break
		}
	}

	return secure

}
