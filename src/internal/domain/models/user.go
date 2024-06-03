package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             uuid.UUID
	Username       Username
	Email          Email
	PasswordHash   Password
	CreationDate   time.Time
	LastUpdateDate time.Time
}

type Username string
type Email string
type Password string

// TODO error handling?
func NewUser(username Username, email Email, password Password) *User {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u := &User{
		ID:             uuid.New(),
		Username:       username,
		Email:          email,
		PasswordHash:   Password(passwordHash),
		CreationDate:   time.Now(),
		LastUpdateDate: time.Now(),
	}
	return u
}
