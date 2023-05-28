package dbConnection

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
var createTableStatus bool
var createTableQuery string
var insertQuery string
var saveUserStatus bool

func DbConnectionVerification(dataSourceName string, createTableQuery string, insertQuery string, firstName string, lastName string, email string, userTickets uint) bool {

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

	_, err = db.Exec(createTableQuery)
	if err != nil {
		fmt.Println("Failed to create table:", err)
		//createTableStatus := true
		//return createTableStatus

	}

	// Prepare the statement
	stmt, err := db.Prepare(insertQuery)
	if err != nil {
		dbValidConnection := false
		return dbValidConnection
	}

	_, err = stmt.Exec(firstName, lastName, email, userTickets)
	fmt.Println("Data inserted successfully!")

	fmt.Println("Table created successfully!")

	return dbValidConnection
}
