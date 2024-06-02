package repositories

import (
	"database/sql"

	"github.com/google/uuid"
	m "github.com/runebone/db-cw-s6-bmstu/internal/domain/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(u *m.User) error {
	query := `
		INSERT INTO user_data
		(id, username, email, pwd_hash, creation_date, last_update_date) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.DB.Exec(query, u.ID, u.Username, u.Email, u.PasswordHash, u.CreationDate, u.LastUpdateDate)
	return err
}

func (r *UserRepository) GetUserByUsername(username m.Username) (*m.User, error) {
	var u *m.User
	query := `
		SELECT
		id, username, email, pwd_hash, creation_date, last_update_date
		FROM user_data
		WHERE username = $1
	`
	row := r.DB.QueryRow(query, username)
	err := row.Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.CreationDate, &u.LastUpdateDate)
	if err != nil {
		return &m.User{}, err
	}
	return u, nil
}

func (r *UserRepository) GetUserByID(id uuid.UUID) (*m.User, error) {
	var u *m.User
	query := `
		SELECT
		id, username, email, pwd_hash, creation_date, last_update_date
		FROM user_data
		WHERE id = $1
	`
	row := r.DB.QueryRow(query, id)
	err := row.Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.CreationDate, &u.LastUpdateDate)
	if err != nil {
		return &m.User{}, err
	}
	return u, nil
}
