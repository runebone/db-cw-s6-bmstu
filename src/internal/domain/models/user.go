package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id             uuid.UUID
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
	password_hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u := &User{
		Id:           uuid.New(),
		Username:     username,
		Email:        email,
		PasswordHash: Password(password_hash),
		CreationDate: time.Now(),
	}
	return u
}
