package models

import (
	"fmt"

	"deepak.com/web_rest/db"
)

type User struct {
	Id       int64
	Email    string `bindings:"required"`
	Password string `bindings:"required"`
}

func (u *User) Save() error {
	
	fmt.Println(u)
	insertQuery := `INSERT INTO USERS(email, password) VALUES(?, ?)`

	stmt, err := db.DB.Prepare(insertQuery)

	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.Email, u.Password)

	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()

	u.Id = id

    defer stmt.Close()
	return nil
}
