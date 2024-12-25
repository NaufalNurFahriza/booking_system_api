package repository

import (
	"database/sql"
	"quiz-sanbercode/structs"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(db *sql.DB, user structs.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	sql := `INSERT INTO users (username, password, created_at, created_by) 
            VALUES ($1, $2, $3, $4)`

	_, err = db.Exec(sql, user.Username, string(hashedPassword),
		time.Now(), user.CreatedBy)
	return err
}

func GetUserByUsername(db *sql.DB, username string) (structs.User, error) {
	var user structs.User
	sql := `SELECT id, username, password, created_at, created_by, 
            modified_at, modified_by FROM users WHERE username = $1`

	err := db.QueryRow(sql, username).Scan(
		&user.ID, &user.Username, &user.Password, &user.CreatedAt,
		&user.CreatedBy, &user.ModifiedAt, &user.ModifiedBy,
	)
	return user, err
}
