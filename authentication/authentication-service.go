package authentication

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/fabruun/go-authentication/contracts"
	"github.com/fabruun/go-authentication/domain"
	"github.com/fabruun/go-authentication/hashing"
	"github.com/fabruun/go-authentication/providers"
	"github.com/fabruun/go-authentication/tokens"
)

type Authentication struct {
	Users *domain.Users
}

func (p *Authentication) Login(email string, password string) *domain.User {
	emailIsValid, passwordIsValid := false, false
	userIsAuthenticated := (!emailIsValid && !passwordIsValid)
	for _, user := range p.Users.Users {
		checkEmail := p.checkEmail(&user, email)
		if checkEmail == nil {
			emailIsValid = true
		}
		checkPassword := p.checkPassword(&user, password)
		if checkPassword == nil {
			passwordIsValid = true
		}
	}
	if userIsAuthenticated {
		user := domain.Get(email)
		return user
	}
	return nil
}

func (p *Authentication) checkEmail(user *domain.User, email string) error {
	if user.Email == email {
		return nil
	}
	format := "No user exists with email: %s"
	message := fmt.Sprintf(format, email)
	log.Printf(message)
	return errors.New(message)
}

func (p *Authentication) checkPassword(user *domain.User, password string) error {
	byteValue := []byte(password)
	passwordHasher := hashing.PasswordHasher{}
	err := passwordHasher.CheckPassword(user, byteValue)
	if err != nil {
		return err
	}
	return nil
}

func (p *Authentication) Register() {
	request := http.Request{}
	jwtTokenGenerator := tokens.JwtTokenGenerator{}
	hashedPassword := p.getHashedPassword(request.Password)
	user := p.createUserObject(&request)
	user.Password = hashedPassword
	jwtTokenGenerator.GenerateToken(&user)
}

func (p *Authentication) getHashedPassword(password string) []byte {
	hasher := hashing.PasswordHasher{}
	hashedPassword := hasher.HashPassword(password)
	if hashedPassword == nil {
		log.Printf("Failed to hash password.")
		return nil
	}
	return hashedPassword
}

func (p *Authentication) createUserObject(request *contracts.CreateNewUser) domain.User {
	user := domain.User{
		Name:      request.Name,
		Email:     request.Email,
		CreatedAt: providers.Now(),
		UpdatedAt: providers.Now(),
	}
	return user
}
