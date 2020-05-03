package models

import (
	"fmt"

	"github.com/lilosir/cyticoffee-api/db/mysql"
	"github.com/lilosir/cyticoffee-api/utils"
)

//Type struct
type Type struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetType will fetch all the typies
func GetType() ([]Type, error) {
	var types []Type
	stmt, err := mysql.DBConn().Prepare("select * from types")
	if err != nil {
		return types, utils.ServerError
	}

	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return types, utils.ServerError
	}

	for rows.Next() {
		t := Type{}
		err := rows.Scan(&t.ID, &t.Name)
		if err != nil {
			return types, utils.ServerError
		}
		types = append(types, t)
	}
	return types, nil
}

// GetOneType will fetch one type
func GetOneType(id string) (Type, error) {
	var onetype Type
	apiErr := utils.ServerError
	stmt, err := mysql.DBConn().Prepare("select * from types where id=? limit = 1")
	if err != nil {
		return onetype, utils.ServerError
	}

	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&onetype.ID, &onetype.Name)
	if err != nil {
		fmt.Println("scan: ", err.Error())
		apiErr.Data = err.Error()
		return onetype, utils.ServerError
	}
	return onetype, nil
}
