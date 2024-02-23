package models

import (
	"errors"

	"example.com/rest-api/database"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashPassword, err := utils.GetHashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}

func (u *User) ValidateCredential() error {
	query := "SELECT * FROM users WHERE email = ? "
	var user User
	row := database.DB.QueryRow(query, u.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		//lint:ignore ST1005 Error string is intentionally in Title case
		return errors.New("Invalid Credentials")
	}
	u.ID = user.ID

	isValidPassword := utils.CheckPassword(u.Password, user.Password)

	if !isValidPassword {
		//lint:ignore ST1005 Error string is intentionally in Title case
		return errors.New("Invalid Credentials")
	}

	return nil
}

func GetUserById(userId int64) (User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	var user User
	rows := database.DB.QueryRow(query, userId)
	err := rows.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil

}
