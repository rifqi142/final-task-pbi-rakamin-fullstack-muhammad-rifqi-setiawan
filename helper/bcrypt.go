package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(p string) string {
	salt := 8
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hash)
}

func CheckPasswordHash(p, hash string) bool {
	password := []byte(p)
	hashedPassword := []byte(hash)

	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	return err == nil
}