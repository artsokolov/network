package auth

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type ErrHashingPassword struct {
	e error
}

func (err *ErrHashingPassword) Error() string {
	return fmt.Sprintf("Error while hashing password: %s", err.e.Error())
}

type ErrComparingPasswords struct {
	e error
}

func (err *ErrComparingPasswords) Error() string {
	return fmt.Sprintf("Error while comparing password: %s", err.e.Error())
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", &ErrHashingPassword{err}
	}

	return string(hash), nil
}

func CheckPasswords(pass, hash string) error {
	e := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if e != nil {
		return &ErrComparingPasswords{e}
	}

	return nil
}
