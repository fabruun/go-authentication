package hashing

import (
	"errors"
	"fmt"
	"log"

	"github.com/fabruun/go-authentication/domain"
	"golang.org/x/crypto/bcrypt"
)

const COST = 11

type PasswordHasher struct{}

func (p *PasswordHasher) HashPassword(password string) []byte {
	bytesValue := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytesValue, COST)
	if err != nil {
		log.Printf("Failed to hash password. Error: %v", err)
		return nil
	}
	return hashedPassword
}

func (p *PasswordHasher) CheckPassword(u *domain.User, password []byte) error {
	err := bcrypt.CompareHashAndPassword(u.Password, password)
	if err != nil {
		message := fmt.Sprintf("Failed to compare passwords. Error: %v", err)
		log.Printf(message)
		return errors.New(message)
	}
	return nil
}
