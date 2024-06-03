package services

import (
	"errors"

	m "github.com/runebone/db-cw-s6-bmstu/internal/domain/models"
	r "github.com/runebone/db-cw-s6-bmstu/internal/domain/repositories"
	"golang.org/x/crypto/bcrypt"
)

// XXX should I use UserRepo *r.UserRepositoryInterface instead like in Go mistake #5?
type UserService struct {
	UserRepo *r.UserRepository
}

func NewUserService(repo *r.UserRepository) *UserService {
	return &UserService{UserRepo: repo}
}

func (s *UserService) RegisterUser(username m.Username, email m.Email, password m.Password) error {
	u := m.NewUser(username, email, password)
	return s.UserRepo.CreateUser(u)
}

// XXX should I return *m.User or m.User?
func (s *UserService) AuthenticateUser(username m.Username, password m.Password) (*m.User, error) {
	usr, err := s.UserRepo.GetUserByUsername(username)
	if err != nil {
		return &m.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(usr.PasswordHash), []byte(password))
	if err != nil {
		return &m.User{}, errors.New("invalid credentials")
	}

	return usr, nil
}
