package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

type Mssql struct {
	dataSource string
	database   string
	windows    bool
	user       string
	pwd        string
	port       int
}

func NewDB() *Mssql {

	m := &Mssql{
		dataSource: "salarsql15.salarynet.local",
		database:   "GlobalSurveyData",
		windows:    false,
		user:       "sa",
		pwd:        "S@lGen2",
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
