package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "khangdev"
	check := "khangde2v"
	hash, _ := hashPassword(password)

	fmt.Println("Password:", password)
	fmt.Println("Hash:", hash)

	match := checkPassword(check, hash)

	fmt.Println("Match:", match)
}
