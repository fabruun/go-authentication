package hashing

import (
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
		log.Fatalf("Failed to hash password. Error: %v", err)
		return nil
	}
	return hashedPassword
}

func (p *PasswordHasher) CheckPassword(u *domain.User, password []byte) {
	err := bcrypt.CompareHashAndPassword(u.Password, password)
	if err != nil {
		log.Fatalf("Failed to compare passwords. Error: %v", err)
		return
	}
}
