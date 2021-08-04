package main

import (
	"ClearCompanyName/utils"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	// var connString string
	// s, port, d, u, p := "192.168.3.211", "1433", "Compensation_Consumer", "sa", "S@lGen2"
	// connString = fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", s, port, d, u, p)

	// //建立连接
	// db, err := sql.Open("mssql", connString)
	// if err != nil {
	// 	log.Fatal("Open Connection failed:", err.Error())
	// }
	// defer db.Close()

	// r := GetResult()
	// count := 0
	// for i := 0; i < len(r); i++ {
	// 	companyID := r[i][0]
	// 	companyName := r[i][1]
	// 	cleanCompanyName := utils.GetDeepCleanCompanyName(companyName)
	// 	// sql := fmt.Sprintf("update tbl_ParentCompany_US_Information set CleanCompanyName = '%s' where TRDCompanyID = '%s'", cleanCompanyName, companyID)
	// 	_, err := db.Exec(`update tbl_ParentCompany_US_Information set CleanCompanyName = ? where TRDCompanyID = ?`, cleanCompanyName, companyID)
	// 	if err != nil {
	// 		fmt.Println("update error! ", companyID, cleanCompanyName)
	// 		return
	// 	}
	// 	count += 1
	// }
	// fmt.Println("updated %d rows", count)
	GetJsonResult()
}

func GetResult() [][]string {
	var connString string
	s, port, d, u, p := "192.168.3.211", "1433", "Compensation_Consumer", "sa", "S@lGen2"
	connString = fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", s, port, d, u, p)

	//建立连接
	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
	}
	defer db.Close()

	sqlstr := `select TOP 10 TRDCompanyID,
					CleanCompanyName,
					SEOCompanyName,
					CleanCompanyDomain
			   FROM tbl_ParentCompany_US_Information			
	`
	// fmt.Println("sqlstr:", sqlstr)
	rows, err := db.Query(sqlstr)

	if err != nil {
		log.Fatal("Query failed:", err.Error())
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatalln(err)
	}

	vals := make([][]byte, len(columns))
	scans := make([]interface{}, len(columns))

	for i := range vals {
		scans[i] = &vals[i]

	}

	var results [][]string
	// results = append(results, columns)
	for rows.Next() {
		err = rows.Scan(scans...)
		if err != nil {
			fmt.Println("Failed to scan row", err)
			return results
		}
		row := make([]string, len(columns))
		for i := range vals {
			row[i] = string(vals[i])
		}
		results = append(results, row)
	}
	if len(results) == 1 {
		warning := []string{"no data found!"}
		results = append(results, warning)
	}
	return results
}

func GetJsonResult() {
	var connString string
	s, port, d, u, p := "192.168.3.211", "1433", "Compensation_Consumer", "sa", "S@lGen2"
	connString = fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", s, port, d, u, p)

	//建立连接
	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
	}
	defer db.Close()

	// sqlstr := `select TOP 10 TRDCompanyID,
	// 				CleanCompanyName,
	// 				SEOCompanyName,
	// 				CleanCompanyDomain
	// 		   FROM tbl_ParentCompany_US_Information
	// `
	sqlstr := `	select *
				from tbl_BonusPayable		
	`
	// fmt.Println("sqlstr:", sqlstr)
	rows, err := db.Query(sqlstr)

	if err != nil {
		log.Fatal("Query failed:", err.Error())
	}
	defer rows.Close()

	result, err := utils.SQLToJSON(rows)
	if err != nil {
		fmt.Println("Faild on SQLToJSON:", err)
		return
	}
	fmt.Println("result:", result)
}
