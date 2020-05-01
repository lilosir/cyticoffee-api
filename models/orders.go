package models

import (
	"fmt"
	"strconv"

	"github.com/lilosir/cyticoffee-api/db/mysql"
	"github.com/lilosir/cyticoffee-api/utils"
)

// OrderItem struct
type OrderItem struct {
	TypeID  int64   `json:"type_id"`
	ItemID  int64   `json:"item_id"`
	Options []int   `json:"options"`
	Amount  int     `json:"amount"`
	Price   float32 `json:"price"`
}

// CreateOrders will create a order for the current user
func CreateOrders(order []OrderItem) error {
	apiErr := utils.ServerError
	tx, err := mysql.DBConn().Begin()
	if err != nil {
		apiErr.Data = "transaction begin error"
		return apiErr
	}
	{
		valueStmt := ""
		values := []interface{}{}
		for i := 0; i < len(order); i++ {
			valueStmt += "(?,?,?,?,?)"
			if i != len(order)-1 {
				valueStmt += ","
			}
			options := ""
			for j := 0; j < len(order[i].Options); j++ {
				options += strconv.Itoa(order[i].Options[j])
				if j != len(order[i].Options)-1 {
					options += ","
				}
			}
			values = append(values, order[i].TypeID, order[i].ItemID, options, order[i].Amount, order[i].Price)
		}
		insetStmt := `insert into order_details
		(type_id, item_id, options, amount, price)
		values ` + valueStmt + `;`
		fmt.Println(insetStmt)

		stmt, err := tx.Prepare(insetStmt)
		if err != nil {
			tx.Rollback()
			fmt.Println("prepare: ", err.Error())
			apiErr.Data = err.Error()
			return apiErr
		}
		defer stmt.Close()

		_, err = stmt.Exec(values...)
		if err != nil {
			tx.Rollback()
			fmt.Println("exec: ", err.Error())
			apiErr.Data = err.Error()
			return apiErr
		}

		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			fmt.Println("commit: ", err.Error())
			apiErr.Data = err.Error()
			return apiErr
		}
	}
	return nil
}
