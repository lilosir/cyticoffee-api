package models

import (
	"errors"
	"fmt"

	"github.com/lilosir/cyticoffee-api/db/mysql"
)

// User struct
type User struct {
	Id        int64  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
}

//UserSignup will create a new user
func UserSignup(u *User) (int64, error) {
	stmt, err := mysql.DBConn().Prepare(
		"insert ignore into tbl_user(`firstname`, `lastname`, `email`, `password`, `phone`) values (?,?,?,?,?)")
	if err != nil {
		fmt.Printf(err.Error())
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(u.Firstname, u.Lastname, u.Email, u.Password, u.Phone)
	if err != nil {
		fmt.Printf(err.Error())
		return 0, err
	}

	id, _ := result.LastInsertId()
	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf(err.Error())
		return 0, err
	}

	if affected <= 0 {
		return 0, errors.New("already exists")
	}

	return id, nil
}
