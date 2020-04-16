package models

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/lilosir/cyticoffee-api/db/mysql"
	"github.com/lilosir/cyticoffee-api/utils"
)

// User struct
type User struct {
	ID         int64  `json:"id"`
	Firstname  string `json:"firstname" binding:"required"`
	Lastname   string `json:"lastname" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	SignupAt   string `json:"signup_at"`
	LastActive string `json:"last_active"`
}

//UserSignup will create a new user
func UserSignup(u User) error {
	stmt, err := mysql.DBConn().Prepare(
		"insert ignore into tbl_user(`firstname`, `lastname`, `email`, `password`, `phone`) values (?,?,?,?,?)")
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(u.Firstname, u.Lastname, u.Email, u.Password, u.Phone)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}

	_, err = result.LastInsertId()
	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}

	if affected <= 0 {
		return utils.NewAPIError(http.StatusConflict, "Email already exists", nil)
	}

	return nil
}

// UserLogIn will check if email and password matching, otherwise return an error
func UserLogIn(email, password string) (User, error) {
	var user User
	stmt, err := mysql.DBConn().Prepare("select * from tbl_user where email=? limit 1")
	if err != nil {
		return user, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(email).Scan(
		&user.ID,
		&user.Firstname,
		&user.Lastname,
		&user.Email,
		&user.Password,
		&user.Phone,
		&user.SignupAt,
		&user.LastActive,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			apiErr := utils.NewAPIError(http.StatusNotFound, "Your email does not exist", nil)
			return user, apiErr
		}
		return user, err
	}

	if user.Password != utils.CreateSha1([]byte(password)) {
		return user, utils.PasswordNotMatch
	}

	return user, nil
}
