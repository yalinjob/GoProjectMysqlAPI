package dbConnection

import (
	"GOLNGCOURSE/helper"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var dataSourceName string
var dbValidConnection bool
var createTableStatus bool
var createTableQuery string
var insertQuery string
var saveUserStatus bool

func DbConnectionVerification(dataSourceName string, userData helper.UserData) bool {

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

	createTableQuery := `
				CREATE TABLE IF NOT EXISTS users (
					id INT AUTO_INCREMENT PRIMARY KEY,
					first_name VARCHAR(50),
					last_name VARCHAR(50),
					email VARCHAR(100),
					number_of_ticket INT

				) `

	_, err = db.Exec(createTableQuery)
	if err != nil {
		fmt.Println("Failed to create table:", err)
		//createTableStatus := true
		//return createTableStatus

	}
	fmt.Println("Data inserted successfullyfirst point!")
	// Prepare the SQL query
	//insertQuery := ` INSERT INTO users (firstName, lastName, email,userTickets) VALUES (?, ?, ?, ?)`

	query := "INSERT INTO users (first_name, last_name, email, number_of_ticket) VALUES (?, ?, ?, ?)"

	// Prepare the statement
	stmt, err := db.Prepare(query)
	if err != nil {
		dbValidConnection := false
		return dbValidConnection

	}

	defer stmt.Close()

	fmt.Println("Data inserted successfullySecond point!")
	//_, err = stmt.Exec(firstName, lastName, email, userTickets)

	_, err = stmt.Exec(userData.FirstName, userData.LastName, userData.Email, userData.NumberOfTicket)
	fmt.Println("Data inserted successfullyLast!")

	fmt.Println("Table created successfully!")

	return dbValidConnection

}
