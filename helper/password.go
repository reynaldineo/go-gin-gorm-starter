package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}

func CompareHashPassword(hashPassword string, plainPassword []byte) (bool, error) {
	hashPass := []byte(hashPassword)
	if err := bcrypt.CompareHashAndPassword(hashPass, plainPassword); err != nil {
		return false, err
	}

	return true, nil
}
