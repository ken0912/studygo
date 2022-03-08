package main

import (
	"database/sql"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	h      bool
	s, S   string
	port   int
	u, U   string
	p, P   string
	d, D   string
	t      string
	fp     string
	number int
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&s, "s", "", "db server name ")
	flag.StringVar(&s, "S", "", "db server name ")
	flag.IntVar(&port, "port", 1433, "db port")
	flag.StringVar(&u, "u", "", "db login user name")
	flag.StringVar(&u, "U", "", "db login user name")
	flag.StringVar(&p, "p", "", "db login password")
	flag.StringVar(&p, "P", "", "db login password")
	flag.StringVar(&d, "d", "", "database name")
	flag.StringVar(&d, "D", "", "database name")
	flag.StringVar(&t, "t", "", "The table in which data needs to be inserted")
	flag.StringVar(&fp, "fp", "", "full path of the csv file needs to be imported")
	flag.IntVar(&number, "number", 5000, "number of per batch")
}

var rowChan = make(chan []interface{}, 1024)

func main() {
	flag.Parse()
	validation()
	if h {
		flag.Usage()
		return
	}
	stop := make(chan bool)
	go Extractdata()
	go WriteData(stop)
	for {
		select {
		case <-stop:
			fmt.Println("Import Done:", time.Now().Format("2006/01/02 15:04:05"))
			return
		default:
			time.Sleep(1 * time.Second)
		}
	}

}

func validation() {
	if fp == "" || t == "" {
		panic("-fp(filepath) or -t(tablename) are not allowed to be empty")
	}

}
func GenerateSqlInsertStr(tablename string, columnlength int) (sqlstr string) {
	sqlstr = `INSERT INTO ` + tablename + ` VALUES`
	//placeholder
	valueslice := make([]string, columnlength)
	for i := 0; i < columnlength; i++ {
		valueslice[i] = "?"
	}
	return sqlstr + "(" + strings.Join(valueslice, ",") + ")"
}

//Extract data from csv file
func Extractdata() {
	f, err := os.Open(fp)
	if err != nil {
		log.Fatalln("os.Open error:", err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.LazyQuotes = true
	id := 0

	columnlen := 0

	for {
		id++
		record, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		if id == 1 {
			columnlen = len(record)
			continue
		}
		recorditf := make([]interface{}, columnlen)
		for i, v := range record {
			recorditf[i] = v
		}
		rowChan <- recorditf
	}
	close(rowChan)
}

//write data to table
func WriteData(stop chan bool) {
	fmt.Println("Start Importing:", time.Now().Format("2006/01/02 15:04:05"))

	//DB建立连接
	connString := fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", s, port, d, u, p)
	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
	}
	defer db.Close()
	//begin tran
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("db.Begin() err:", err)
	}
	//get column title
	firstrow := <-rowChan
	sqlstr := GenerateSqlInsertStr(t, len(firstrow))
	// fmt.Println("sqlstr:", sqlstr)
	//write data
	RowAffected := 0
	ErrorCount := 0
	for row := range rowChan {
		// _, err := DBExec(db, sqlstr, row)
		_, err := tx.Exec(sqlstr, row...)
		if err != nil {
			log.Fatalln("DBExec() err:", err)
			ErrorCount += 1
		} else {
			RowAffected += 1
		}
		// fmt.Printf("\r Row:%d", RowAffected)
	}
	if ErrorCount == 0 {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	fmt.Println("")
	fmt.Println("RowAffected:", RowAffected)
	fmt.Println("ErrorCount:", ErrorCount)
	stop <- true
}
func DBExec(db *sql.DB, stmt string, valueArgs []interface{}) (rowaffected int64, err error) {
	r, err := db.Exec(stmt, valueArgs...)

	if err != nil {
		fmt.Println("db.Exec() err:", err)

	}
	rowaffected, err = r.RowsAffected()
	if err != nil {
		log.Fatalln("r.RowsAffected() err:", err)
	}
	return
}
func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
func BoolToBit(s string) int {
	if s == "TRUE" {
		return 1
	} else {
		return 0
	}
}
