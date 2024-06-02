package repositories

import (
	"database/sql"

	"github.com/runebone/db-cw-s6-bmstu/internal/domain/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(u models.User) error {
	query := `
		INSERT INTO user_data
		(id, username, email, pwd_hash, creation_date, last_update_date) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.DB.Exec(query, u.Id, u.Username, u.Email, u.PasswordHash, u.CreationDate, u.LastUpdateDate)
	return err
}

func (r *UserRepository) GetUserByUsername(username models.Username) (models.User, error) {
	var u models.User
	query := `
		SELECT
		id, username, email, pwd_hash, creation_date, last_update_date
		FROM user_data
		WHERE username = $1
	`
	row := r.DB.QueryRow(query, username)
	err := row.Scan(&u.Id, &u.Username, &u.Email, &u.PasswordHash, &u.CreationDate, &u.LastUpdateDate)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}
