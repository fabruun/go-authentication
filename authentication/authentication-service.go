package authentication

import (
	"log"

	"github.com/fabruun/go-authentication/contracts"
	"github.com/fabruun/go-authentication/domain"
	"github.com/fabruun/go-authentication/hashing"
	"github.com/fabruun/go-authentication/providers"
	"github.com/fabruun/go-authentication/tokens"
)

type Authentication struct {
	Users *domain.Users
}

func (p *Authentication) Login(email string, password string) (user *domain.User) {
	for _, user := range p.Users.Users {
		if user.Email == email {
			return &user
		}
	}
	return nil
}

func (p *Authentication) Register(request contracts.CreateNewUser) {
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
		log.Fatalf("Failed to hash password.")
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
