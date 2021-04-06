// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

type Post struct {
	UserId     int64  `json:"userid"`
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	Testcolumn string `json:"testcolumn"`
}

func main() {
	var isdebug = true
	var server = "salarysql15.salarynet.local"
	var port = 1433
	var user = "sa"
	var password = "S@lGen2"
	var database = "Pubs"

	//连接字符串
	connString := fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", server, port, database, user, password)
	if isdebug {
		fmt.Println(connString)
	}
	//建立连接
	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
	}
	defer db.Close()

	//通过连接对象执行查询
	rows, err := db.Query(`select TOP 1 * from tbl_IncumProfile_Export_20200623`)
	if err != nil {
		log.Fatal("Query failed:", err.Error())
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("Failed to get columns", err)
		return
	}

	// Result is your slice string.
	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))

	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	for i, _ := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			fmt.Println("Failed to scan row", err)
			return
		}

		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\N"
			} else {
				result[i] = string(raw)
			}
		}

		fmt.Printf("%#v\n", result)
	}
}
