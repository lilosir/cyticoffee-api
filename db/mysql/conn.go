package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/cyticoffee")
	if err != nil {
		fmt.Println("!!!")
		panic(err.Error())
	}

	db.SetMaxOpenConns(1000)
	err = db.Ping()
	if err != nil {
		fmt.Printf("cannot open connection with mysql, %s\n", err.Error())
	}

}

// DBConn returns database connection object
func DBConn() *sql.DB {
	return db
}
