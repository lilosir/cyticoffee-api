package models

import (
	"fmt"
	"strings"

	"github.com/lilosir/cyticoffee-api/db/mysql"
	"github.com/lilosir/cyticoffee-api/utils"
)

//Options struct
type Options struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

// GetOptions returns all the types by id
func GetOptions(ids string) ([]Options, error) {
	apiErr := utils.ServerError

	whereIds := ""
	for _, o := range strings.Split(ids, ",") {
		whereIds += o + ","
	}
	whereIds = strings.TrimSuffix(whereIds, ",")
	stmt, err := mysql.DBConn().Prepare("select * from options where id in (" + whereIds + ")")
	if err != nil {
		fmt.Println("prepare select from options ", err.Error())
		apiErr.Data = err.Error()
		return nil, apiErr
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("query: ", err.Error())
		apiErr.Data = err.Error()
		return nil, apiErr
	}

	var options []Options
	for rows.Next() {
		option := Options{}
		err := rows.Scan(&option.ID, &option.Name, &option.Type)
		if err != nil {
			fmt.Println("scan: ", err.Error())
			apiErr.Data = err.Error()
			return nil, apiErr
		}
		options = append(options, option)
	}
	return options, nil
}

// GetAllOptions returns all the options
func GetAllOptions() ([]Options, error) {
	apiErr := utils.ServerError

	stmt, err := mysql.DBConn().Prepare("select * from options")
	if err != nil {
		fmt.Println("prepare: ", err.Error())
		apiErr.Data = err.Error()
		return nil, apiErr
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("query: ", err.Error())
		apiErr.Data = err.Error()
		return nil, apiErr
	}

	var options []Options
	for rows.Next() {
		option := Options{}
		err := rows.Scan(&option.ID, &option.Name, &option.Type)
		if err != nil {
			fmt.Println("scan: ", err.Error())
			apiErr.Data = err.Error()
			return nil, apiErr
		}
		options = append(options, option)
	}
	return options, nil
}
