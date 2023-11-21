package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func ComparePassword(password string, confirmPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(confirmPassword))
	if err != nil {
		return false, err
	}

	return true, err
}