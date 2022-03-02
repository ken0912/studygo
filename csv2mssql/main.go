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
	h    bool
	s, S string
	port int
	u, U string
	p, P string
	d, D string
	t    string
	fp   string
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
func GenerateSqlInsertStr(tablename string, columnlength int) (stmt, placeholder string) {
	stmt = `INSERT INTO ` + tablename + ` VALUES`
	// for i := 0; i < columnlength; i++ {
	// 	columnfmt += `?,`
	// }
	// columnfmt += `)`
	// columnfmt = strings.Replace(columnfmt, ",)", ")", -1)

	//placeholder
	valueslice := make([]string, columnlength)
	for i := 0; i < columnlength; i++ {
		valueslice[i] = "?"
	}
	return stmt, "(" + strings.Join(valueslice, ",") + ")"
}

//Extract data from csv file
func Extractdata() {
	f, err := os.Open(fp)
	if err != nil {
		log.Fatalln("os.Open error:", err)
	}
	defer f.Close()
	r := csv.NewReader(f)
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
	//get column title
	firstrow := <-rowChan
	stmt, placeholder := GenerateSqlInsertStr(t, len(firstrow))
	//write data
	// 存放 (?, ?) 的slice
	valueStrings := make([]string, 0, 1024*2)
	// 存放值的slice
	valueArgs := make([]interface{}, 0, 1024*2)
	for row := range rowChan {
		valueStrings = append(valueStrings, placeholder)
		valueArgs = append(valueArgs, row...)
	}
	stmt = fmt.Sprintf("%s %s", stmt, strings.Join(valueStrings, ","))
	r, err := db.Exec(stmt, valueArgs...)
	if err != nil {
		fmt.Println("err:", err)

	}
	rowaffected, err := r.RowsAffected()
	if err != nil {
		log.Fatalln("r.RowsAffected() err:", err)
	}
	fmt.Println("rowaffected:", rowaffected)
	stop <- true
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
