package domain

import (
	"crypto"
	"time"

	"github.com/google/uuid"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	ID        uuid.UUID   `json:"id"`
	Name      string      `json:"name"`
	Email     string      `json:"email"`
	Password  crypto.Hash `json:"password"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Token     string      `json:"token"`
}
