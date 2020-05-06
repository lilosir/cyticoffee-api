package models

import (
	"fmt"
	"strings"

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
	orderDetailsIdsIntArray := make([]int, len(order))
	totalPrice := float32(0)
	{
		for i := 0; i < len(order); i++ {
			stmt, err := tx.Prepare("insert into order_details " +
				"(type_id, item_id, options, amount, price)" +
				"values (?,?,?,?,?)")
			if err != nil {
				tx.Rollback()
				fmt.Println("prepare insert into order_details: ", err.Error())
				apiErr.Data = err.Error()
				return apiErr
			}
			defer stmt.Close()

			stringOptions := utils.IntArrayToString(order[i].Options)
			result, err := stmt.Exec(order[i].TypeID, order[i].ItemID, stringOptions, order[i].Amount, order[i].Price)
			if err != nil {
				tx.Rollback()
				fmt.Println("Exec insert into order_details: ", err.Error())
				apiErr.Data = err.Error()
				return apiErr
			}
			orderDetailID, err := result.LastInsertId()
			if err != nil {
				tx.Rollback()
				fmt.Println("insert into order_details return id: ", err.Error())
				apiErr.Data = err.Error()
				return apiErr
			}
			totalPrice += float32(order[i].Amount) * order[i].Price
			orderDetailsIdsIntArray[i] = int(orderDetailID)
		}
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

		orderDetailsIdsString := utils.IntArrayToString(orderDetailsIdsIntArray)
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
	defer stmt.Close()

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

// Order struct
type Order struct {
	OrderDetail []OrderDetail `json:"order_detail"`
	TotalPrice  float32       `json:"total_price"`
	Status      string        `json:"status"`
	CreatedAt   string        `json:"create_at"`
}

// OrderDetail struct
type OrderDetail struct {
	Type    string    `json:"type"`
	Options []Options `json:"options"`
	Amount  int       `json:"amount"`
	Price   float32   `json:"price"`
}

// TempOrder struct
type TempOrder struct {
	OrderDetailIDs string `json:"order_detail_ids"`
}

// GetOrderDetails return an order history details
func GetOrderDetails(orderID string) (Order, error) {
	apiErr := utils.ServerError
	var order Order
	orderDetailIDs := ""
	{
		stmt, err := mysql.DBConn().Prepare(`
			select o.order_detail_ids, o.total_price, s.status, o.create_at 
			from orders o
			left join status s
			on o.status = s.id
			where o.id=?
		`)
		if err != nil {
			fmt.Println("prepare select order detail ids: ", err.Error())
			apiErr.Data = err.Error()
			return order, apiErr
		}
		defer stmt.Close()

		err = stmt.QueryRow(orderID).Scan(&orderDetailIDs, &order.TotalPrice, &order.Status, &order.CreatedAt)
		if err != nil {
			fmt.Println("scan order detail ids: ", err.Error())
			apiErr.Data = err.Error()
			return order, apiErr
		}
	}
	// search order details
	{
		deIds := strings.Split(orderDetailIDs, ",")
		queryOrderInterface := make([]interface{}, len(deIds))
		preS := ""
		for i := 0; i < len(deIds); i++ {
			preS += "?,"
			queryOrderInterface[i] = deIds[i]
		}
		preS = strings.TrimSuffix(preS, ",")
		q := "select t.name, o.item_id, o.options, o.amount, o.price " +
			"from order_details o " +
			"left join types t " +
			"on t.id = o.type_id " +
			"where o.id in (" + preS + ")"
		stmt, err := mysql.DBConn().Prepare(q)
		if err != nil {
			fmt.Println("prepare select order details", err.Error())
			apiErr.Data = err.Error()
			return order, apiErr
		}
		defer stmt.Close()

		rows, err := stmt.Query(queryOrderInterface...)
		if err != nil {
			fmt.Println("prepare query order details", err.Error())
			apiErr.Data = err.Error()
			return order, apiErr
		}
		for rows.Next() {
			orderDetail := OrderDetail{}
			itemID := 0
			options := ""
			err = rows.Scan(&orderDetail.Type, &itemID, &options, &orderDetail.Amount, &orderDetail.Price)
			if err != nil {
				fmt.Println("scan order details", err.Error())
				apiErr.Data = err.Error()
				return order, apiErr
			}
			gd, err := GetGoods(itemID, orderDetail.Type)
			if err != nil {
				return order, err
			}
			orderDetail.Type = gd.Name
			fmt.Println(itemID, options, orderDetail)

			orderDetail.Options, err = GetOptions(options)
			if err != nil {
				return order, err
			}

			order.OrderDetail = append(order.OrderDetail, orderDetail)
		}
	}

	return order, nil
}
