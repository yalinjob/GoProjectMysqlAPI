package dbConfig

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var dbUser string
var doPass string
var dbHost string
var dbName string
var dbPort string
var dataSourceName string
var dbValidConnection bool

// Create a MySQL database connection string

//dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

// Open a connection to the MySQL database
func dbConnectionVerification(dataSourceName string) bool {

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Attempt to ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to the MySQL database!")

	dbValidConnection := true
	return dbValidConnection
}
