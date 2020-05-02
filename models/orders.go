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

// Orders struct
type Orders struct {
	ID             int64   `json:"id"`
	OrderDetailIDs string  `json:"order_detail_ids"`
	TotalPrice     float32 `json:"total_price"`
	Status         int     `json:"status"`
	CreatedAt      string  `json:"create_at"`
}

// CreateOrders will create a order for the current user
func CreateOrders(order []OrderItem, userID interface{}) error {
	apiErr := utils.ServerError
	tx, err := mysql.DBConn().Begin()
	if err != nil {
		apiErr.Data = "transaction begin error"
		return apiErr
	}
	orderDetailsIdsString := ""
	totalPrice := float32(0)
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

			//add to total price
			totalPrice += float32(order[i].Amount) * order[i].Price
		}
		insetStmt := `insert into order_details
		(type_id, item_id, options, amount, price)
		values ` + valueStmt + `;`
		// fmt.Println(insetStmt)

		stmt, err := tx.Prepare(insetStmt)
		if err != nil {
			tx.Rollback()
			fmt.Println("prepare: ", err.Error())
			apiErr.Data = err.Error()
			return apiErr
		}
		defer stmt.Close()

		result, err := stmt.Exec(values...)
		if err != nil {
			tx.Rollback()
			fmt.Println("exec: ", err.Error())
			apiErr.Data = err.Error()
			return apiErr
		}

		firstID, _ := result.LastInsertId()
		rows, _ := result.RowsAffected()

		generatedOrderDetailsIds := make([]int64, rows)
		for i := 0; i < int(rows); i++ {
			generatedOrderDetailsIds[i] = firstID + int64(i)
		}
		orderDetailsIdsString = utils.Int64ArrayToString(generatedOrderDetailsIds)
		// fmt.Println(orderDetailsIdsString)
	}
	{
		stmt, err := tx.Prepare("Insert into orders (user_id, order_detail_ids, total_price, status) values (?,?,?,?)")
		if err != nil {
			tx.Rollback()
			fmt.Println("Prepare: ", err.Error())
			apiErr.Data = err.Error()
			return apiErr
		}
		defer stmt.Close()

		_, err = stmt.Exec(userID, orderDetailsIdsString, totalPrice, 0)
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

// GetMyOrders will return all my orders
func GetMyOrders(userID string) ([]Orders, error) {
	apiErr := utils.ServerError
	var orders []Orders
	stmt, err := mysql.DBConn().Prepare("select id, order_detail_ids, total_price, status, create_at from orders where user_id = ?")
	if err != nil {
		fmt.Println("prepare: ", err.Error())
		apiErr.Data = err.Error()
		return orders, apiErr
	}

	rows, err := stmt.Query(userID)
	if err != nil {
		fmt.Println("query: ", err.Error())
		apiErr.Data = err.Error()
		return orders, apiErr
	}

	for rows.Next() {
		o := Orders{}
		err = rows.Scan(
			&o.ID,
			&o.OrderDetailIDs,
			&o.TotalPrice,
			&o.Status,
			&o.CreatedAt,
		)
		if err != nil {
			fmt.Println("scan: ", err.Error())
			apiErr.Data = err.Error()
			return orders, apiErr
		}
		orders = append(orders, o)
	}

	return orders, nil
}
