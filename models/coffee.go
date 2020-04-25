package models

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/lilosir/cyticoffee-api/db/mysql"
	"github.com/lilosir/cyticoffee-api/utils"
)

// CoffeeBrief struct
type CoffeeBrief struct {
	ID      int         `json:"id"`
	Name    string      `json:"name"`
	Image   interface{} `json:"image"`
	Feature int         `json:"feature"`
}

// Coffee struct
type Coffee struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	BasePrice float32     `json:"base_price"`
	Image     interface{} `json:"image"`
	Options   string      `json:"options"`
	Feature   int         `json:"feature"`
}

// CoffeeDetail struct
type CoffeeDetail struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	BasePrice float32     `json:"base_price"`
	Image     interface{} `json:"image"`
	Options   []Options   `json:"options"`
	Feature   int         `json:"feature"`
}

//GetAllCoffee returns all the coffee
func GetAllCoffee() ([]CoffeeBrief, error) {
	var allCoffee []CoffeeBrief
	apiErr := utils.ServerError
	stmt, err := mysql.DBConn().Prepare("select id, name, image, feature from coffee")
	if err != nil {
		fmt.Println("prepare: ", err)
		apiErr.Data = err.Error()
		return allCoffee, apiErr
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("query: ", err)
		apiErr.Data = err.Error()
		return allCoffee, apiErr
	}

	for rows.Next() {
		coffee := CoffeeBrief{}
		err = rows.Scan(
			&coffee.ID,
			&coffee.Name,
			&coffee.Image,
			&coffee.Feature,
		)
		if err != nil {
			fmt.Println("scan: ", err.Error())
			apiErr.Data = err.Error()
			return allCoffee, apiErr
		}
		allCoffee = append(allCoffee, coffee)
	}

	return allCoffee, nil
}

// GetCoffee return one type of coffee query with id
func GetCoffee(id int) (CoffeeDetail, error) {
	var coffee Coffee
	var coffeeDetail CoffeeDetail
	apiErr := utils.ServerError
	stmt, err := mysql.DBConn().Prepare("select * from coffee where id=? limit 1")
	if err != nil {
		fmt.Println("prepare: ", err.Error())
		apiErr.Data = err.Error()
		return coffeeDetail, apiErr
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(
		&coffee.ID,
		&coffee.Name,
		&coffee.BasePrice,
		&coffee.Image,
		&coffee.Options,
		&coffee.Feature,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return coffeeDetail, utils.NotFound
		}
		fmt.Println("query and scan: ", err.Error())
		apiErr.Data = err.Error()
		return coffeeDetail, apiErr
	}

	coffeeDetail.ID = coffee.ID
	coffeeDetail.Name = coffee.Name
	coffeeDetail.BasePrice = coffee.BasePrice
	coffeeDetail.Image = coffee.Image
	coffeeDetail.Feature = coffee.Feature

	whereIds := ""
	for _, o := range strings.Split(coffee.Options, ",") {
		whereIds += o + ","
	}
	whereIds = strings.TrimSuffix(whereIds, ",")
	stmt, err = mysql.DBConn().Prepare("select * from options where id in (" + whereIds + ")")
	if err != nil {
		fmt.Println("prepare: ", err.Error())
		apiErr.Data = err.Error()
		return coffeeDetail, apiErr
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("query: ", err.Error())
		apiErr.Data = err.Error()
		return coffeeDetail, apiErr
	}

	var options []Options
	for rows.Next() {
		option := Options{}
		err := rows.Scan(&option.ID, &option.Name, &option.Type)
		if err != nil {
			fmt.Println("scan: ", err.Error())
			apiErr.Data = err.Error()
			return coffeeDetail, apiErr
		}
		options = append(options, option)
	}

	coffeeDetail.Options = options
	return coffeeDetail, nil
}
