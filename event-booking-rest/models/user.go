package models

import (
	"errors"

	"example.com/event-booking/db"
	"example.com/event-booking/utils"
)

type UserSignin struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var hasher = utils.NewArgon2id()

func (u *User) Save() error {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := hasher.Hash([]byte(u.Password), nil)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Name, u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId

	return err
}

func (u *UserSignin) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var storedPassword string
	err := row.Scan(&u.ID, &storedPassword)
	if err != nil {
		return errors.New("invalid credentials")
	}

	err = hasher.Compare(storedPassword, []byte(u.Password))
	if err != nil {
		return err
	}

	return nil
}
