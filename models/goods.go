package models

import (
	"database/sql"
	"fmt"

	"github.com/lilosir/cyticoffee-api/db/mysql"
	"github.com/lilosir/cyticoffee-api/utils"
)

// GoodsBrief struct
type GoodsBrief struct {
	ID      int         `json:"id"`
	Name    string      `json:"name"`
	Image   interface{} `json:"image"`
	Feature int         `json:"feature"`
}

// Goods struct
type Goods struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	BasePrice float32     `json:"base_price"`
	Image     interface{} `json:"image"`
	Options   string      `json:"options"`
	Feature   int         `json:"feature"`
}

// GoodsDetail struct
type GoodsDetail struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	BasePrice float32     `json:"base_price"`
	Image     interface{} `json:"image"`
	Options   []Options   `json:"options"`
	Feature   int         `json:"feature"`
}

// GetAllGoods will reaturn all the good according from differenct tables
func GetAllGoods(table string) ([]GoodsBrief, error) {
	var allGoods []GoodsBrief
	apiErr := utils.ServerError
	stmt, err := mysql.DBConn().Prepare("select id, name, image, feature from " + table)
	if err != nil {
		fmt.Println("prepare: ", err)
		apiErr.Data = err.Error()
		return allGoods, apiErr
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("query: ", err)
		apiErr.Data = err.Error()
		return allGoods, apiErr
	}

	for rows.Next() {
		goods := GoodsBrief{}
		err = rows.Scan(
			&goods.ID,
			&goods.Name,
			&goods.Image,
			&goods.Feature,
		)
		if err != nil {
			fmt.Println("scan: ", err.Error())
			apiErr.Data = err.Error()
			return allGoods, apiErr
		}
		allGoods = append(allGoods, goods)
	}

	return allGoods, nil
}

// GetGoods return one type of goods query with id
func GetGoods(id int, table string) (GoodsDetail, error) {
	var goods Goods
	var goodsDetail GoodsDetail

	apiErr := utils.ServerError
	stmt, err := mysql.DBConn().Prepare("select * from " + table + " where id=? limit 1")
	if err != nil {
		fmt.Println("prepare: ", err.Error())
		apiErr.Data = err.Error()
		return goodsDetail, apiErr
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(
		&goods.ID,
		&goods.Name,
		&goods.BasePrice,
		&goods.Image,
		&goods.Options,
		&goods.Feature,
	)
	if err != nil {
		fmt.Println("query and scan: ", err.Error())
		if err == sql.ErrNoRows {
			return goodsDetail, nil
		}
		apiErr.Data = err.Error()
		return goodsDetail, apiErr
	}

	goodsDetail.ID = goods.ID
	goodsDetail.Name = goods.Name
	goodsDetail.BasePrice = goods.BasePrice
	goodsDetail.Image = goods.Image
	goodsDetail.Feature = goods.Feature

	if goods.Options == "" {
		goodsDetail.Options = []Options{}
	} else {
		goodsDetail.Options, err = GetOptions(goods.Options)
		if err != nil {
			return goodsDetail, err
		}
	}

	return goodsDetail, nil
}
