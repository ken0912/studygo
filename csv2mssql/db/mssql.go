package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

// Mssql ...
type Mssql struct {
	dataSource string
	database   string
	windows    bool
	user       string
	pwd        string
	port       int
}

// NewDB initdb
func NewMssqlClient() *Mssql {
	//config file path
	// var path = "././config.ini"
	// cfg, err := ini.Load(path)
	// if err != nil {
	// 	fmt.Printf("Fail to read file: %v", err)
	// 	os.Exit(1)
	// }
	// server := cfg.Section("ipas_db").Key("server").String()
	// database := cfg.Section("ipas_db").Key("database").String()
	// user := cfg.Section("ipas_db").Key("user").String()
	// pwd := cfg.Section("ipas_db").Key("pwd").String()
	// port := cfg.Section("ipas_db").Key("port").MustInt()
	m := &Mssql{
		dataSource: `PTPC-39PWGQ2\SQL2016`,
		database:   "DB1",
		windows:    false,
		user:       "sa",
		pwd:        "salary.com",
		port:       1433,
	}

	return m
}

//Open database
func (m *Mssql) Open() (db *sql.DB, err error) {
	connString := fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", m.dataSource, m.port, m.database, m.user, m.pwd)
	db, err = sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
		return nil, err

	}
	return db, nil
}
