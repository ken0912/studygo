package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"mongodataload/database"
	"sync"
	"time"

	"github.com/demdxx/gocast"
	_ "github.com/denisenkom/go-mssqldb"
)

func init() {
	// database.MustConnect("mongodb://Salarysql38.salarynet.local:27017", "test")
	database.MustConnect("mongodb://Salarysqldev11.salarynet.local:27019", "GlobalSurveyData")
}

func main() {
	// value := gomongo.Instance.FindOne("t", bson.M{
	// 	"age": 18,
	// })
	// fmt.Println(value.age)
	// fmt.Printf("%v", value)
	var connString string
	s, port, d, u, p := "192.168.3.153", "1433", "GlobalSurveyData", "sa", "S@lGen2"
	// s, port, d, u, p := "192.168.3.211", "1433", "JobBoardTemp", "sa", "S@lGen2"
	connString = fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", s, port, d, u, p)

	//建立连接
	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
	}
	defer db.Close()

	sqlquery := `
		SELECT TOP 10000 * 
		FROM tbl_IncumProfile 
		where SurveyCode = 'IPAS' 
			and companycode = 'D36' 
		order by LoadDate
		`
	getDataFromIPASStart := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("getDataStarts:", getDataFromIPASStart)

	// err = GetIncumProfileToMapToMongoDB(db, sqlquery)
	// if err != nil {
	// 	log.Fatalln("err:", err)
	// }
	result, err := GetIncumProfileToMap(db, sqlquery)

	getDataFromIPASEnd := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("getDataFromIPASEnd", getDataFromIPASEnd)

	insertDataStarts := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("insertDataStarts:", insertDataStarts)

	affectCount := database.Instance.InsertMany("tbl_IncumProfile2", result)

	insertDataEnd := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("insertDataEnd:", insertDataEnd)
	fmt.Println("affectCount", affectCount)

}

func GetIncumProfileToMapToMongoDB(db *sql.DB, sqlstr string, args ...interface{}) error {
	fmt.Println("querystarts", time.Now().Format("2006-01-02 15:04:05"))
	rows, err := db.Query(sqlstr, args...)
	fmt.Println("queryend", time.Now().Format("2006-01-02 15:04:05"))
	columns, err := rows.Columns()
	if err != nil {
		return err
	}
	count := len(columns)
	// tableData := make([]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	num := 1
	var wg sync.WaitGroup
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			// var v interface{}
			val := values[i]
			// entry[col] = convertRow(val)
			entry[col] = val
		}
		/*
			IncumbentID := entry["IncumbentID"].(string)
			//get data from tbl_IncumNumber
			sqlstr1 := `exec usp_GetIncumNumber @SurveyCode = 'IPAS',@IncumbentID = ?`
			// fmt.Println("querystarts:", time.Now().Format("2006-01-02 15:04:05"))
			rows, err := db.Query(sqlstr1, IncumbentID)
			// fmt.Println("queryend:", time.Now().Format("2006-01-02 15:04:05"))
			if err != nil {
				fmt.Println("Query error:", err)
				return nil, err
			}
			res_IncumNumber, _ := GetIncumDetail(rows)
			entry["IncumbentNumber"] = res_IncumNumber

			//get data from tbl_IncumFlag
			sqlstr1 = `exec usp_GetIncumFlag @SurveyCode = 'IPAS',@IncumbentID = ?`
			// fmt.Println("query1starts:", time.Now().Format("2006-01-02 15:04:05"))
			rows, err = db.Query(sqlstr1, IncumbentID)
			// fmt.Println("query2end:", time.Now().Format("2006-01-02 15:04:05"))
			if err != nil {
				fmt.Println("Query error:", err)
				return nil, err
			}
			res_IncumFlag, _ := GetIncumDetail(rows)
			entry["IncumbentFlag"] = res_IncumFlag
		*/
		// tableData = append(tableData, entry)
		// database.InsertOne("tbl_JobPosting", entry)

		wg.Add(1)
		go func() {

			database.Instance.InsertOne("tbl_IncumProfile1", entry)
			wg.Done()
		}()

		fmt.Println("num:", num)
		num += 1

	}
	wg.Wait()
	return nil
}

func GetIncumProfileToMap(db *sql.DB, sqlstr string, args ...interface{}) ([]interface{}, error) {
	fmt.Println("querystarts", time.Now().Format("2006-01-02 15:04:05"))
	rows, err := db.Query(sqlstr, args...)
	fmt.Println("queryend", time.Now().Format("2006-01-02 15:04:05"))
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	count := len(columns)
	tableData := make([]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			// var v interface{}
			val := values[i]
			// entry[col] = convertRow(val)
			entry[col] = val
		}
		/*
			IncumbentID := entry["IncumbentID"].(string)
			//get data from tbl_IncumNumber
			sqlstr1 := `exec usp_GetIncumNumber @SurveyCode = 'IPAS',@IncumbentID = ?`
			// fmt.Println("querystarts:", time.Now().Format("2006-01-02 15:04:05"))
			rows, err := db.Query(sqlstr1, IncumbentID)
			// fmt.Println("queryend:", time.Now().Format("2006-01-02 15:04:05"))
			if err != nil {
				fmt.Println("Query error:", err)
				return nil, err
			}
			res_IncumNumber, _ := GetIncumDetail(rows)
			entry["IncumbentNumber"] = res_IncumNumber

			//get data from tbl_IncumFlag
			sqlstr1 = `exec usp_GetIncumFlag @SurveyCode = 'IPAS',@IncumbentID = ?`
			// fmt.Println("query1starts:", time.Now().Format("2006-01-02 15:04:05"))
			rows, err = db.Query(sqlstr1, IncumbentID)
			// fmt.Println("query2end:", time.Now().Format("2006-01-02 15:04:05"))
			if err != nil {
				fmt.Println("Query error:", err)
				return nil, err
			}
			res_IncumFlag, _ := GetIncumDetail(rows)
			entry["IncumbentFlag"] = res_IncumFlag
		*/
		tableData = append(tableData, entry)

	}

	return tableData, nil
}

func GetIncumDetail(rows *sql.Rows) ([]interface{}, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	count := len(columns)
	tableData := make([]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			// var v interface{}
			val := values[i]
			// entry[col] = convertRow(val)
			entry[col] = val
		}

		tableData = append(tableData, entry)
	}

	return tableData, nil
}

func SQLToJSON(rows *sql.Rows) (string, error) {
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			// var v interface{}
			val := values[i]

			entry[col] = convertRow(val)
		}
		// fmt.Println("entry:", entry)
		tableData = append(tableData, entry)
	}
	byteBuf := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(byteBuf)
	encoder.SetEscapeHTML(false)
	err = encoder.Encode(tableData)
	if err != nil {
		panic(err)
	}

	return byteBuf.String(), nil
}

func convertRow(row interface{}) interface{} {
	switch row.(type) {
	case int:
		return gocast.ToInt(row)
	case string:
		return gocast.ToString(row)
	case []byte:
		return gocast.ToString(row)
	case bool:
		return gocast.ToBool(row)
	case float32:
		return gocast.ToFloat(row)
	case float64:
		return gocast.ToFloat(row)
	}
	return row
}
