package models

import (
	"errors"
	"deepak.com/web_rest/db"
)

type User struct {
	Id       int64
	Email    string `bindings:"required"`
	Password string `bindings:"required"`
}

func (u *User) Save() error {
	
    tx := db.DB.Create(&u)
	return tx.Error
}

func (u *User) Login() error {

	var password string = u.Password

	tx := db.DB.Find(&u, "email = ?", u.Email)

	if tx.RowsAffected == 0 || !(password == u.Password) {
		return errors.New("invalid user credentials")
	}

	return tx.Error

}
