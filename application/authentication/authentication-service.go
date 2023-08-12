package application

import (
	"github.com/fabruun/go-authentication/domain"
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

func (p *Authentication) Register(email string, password string, u *domain.User) {
}
