package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var username, password, hostname, dbname string

func initConfMySQL() {
	username = "homestead"
	password = "!Secret1234"
	hostname = "localhost:3306"
	dbname = "gin_micro_tm"
}

// DB database global
var DB *sql.DB

// SetupDB db
func SetupDB() (*sql.DB, error) {
	initConfMySQL()

	db, err := sql.Open("mysql", confMysql(dbname))

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func confMysql(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}
