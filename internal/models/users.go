package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	stmt := `INSERT INTO users(Name, Email ,hashed_password,Created)
			VALUES(?, ?,?, UTC_TIMESTAMP());`

	_, err = m.DB.Exec(stmt, name, email, string(hashedPassword))
	if err != nil {
		var mySQLError *mysql.MySQLError
		if errors.As(err, &mySQLError) {
			if mySQLError.Number == 1062 && strings.Contains(mySQLError.Message, "users_uc_email") {
				return ErrDuplicateEmail
			}
		}
		return err
	}

	return nil
}

// func (m *UserModel) Exists(email string) (*Users, error) {
// stmt := `SELECT * FROM users
// 		 WHERE email = ?`
// row := m.DB.QueryRow(stmt, email)
// user := &Users{}
// err := row.Scan(user.Id, user.Email, user.Password)
// if err != nil {
// 	if errors.Is(err, sql.ErrNoRows) {
// 		return nil, ErrNoRecord
// 	} else {
// 		return nil, err
// 	}

// }
// return user, nil
// }

func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}
