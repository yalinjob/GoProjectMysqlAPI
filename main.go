package main

import (
	"GOLNGCOURSE/dbConnection"
	"GOLNGCOURSE/helper"

	"fmt"
)

type listOfAllEmails []listOfAllEmails
type getAllUserList []getAllUserList
type getAllEmailAdd []getAllEmailAdd

var conferenceTickets int = 50
var valideUserExsitingResult bool
var isValidUserExsiting bool = true
var isConnnectionIsValid bool = true
var ifValideUserExsiting string
var amountTicket uint
var remainingTickets uint = 50

var dbConnectionStatus bool
var dbCreationTable bool

// Connection parameters

func main() {

	fmt.Printf("We have total of %v tickets and %v\n are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend")

	for {

		firstName, lastName, email, userTickets := helper.GetUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			//Validation of User Exsiting

			valideUserExsitingResult := helper.ValideUserExsiting(firstName)
			fmt.Printf("Indication if user is exsiting   : %v\n ", valideUserExsitingResult)
			if valideUserExsitingResult == true {

				fmt.Printf("You alreday exsiting in the system\n  ")

			} else {

				//Presenting the Ticket Amount

				amountTicket = helper.BookTicket(userTickets, firstName, lastName, email)

				fmt.Printf("The amount of ticket left  : %v\n ", amountTicket)

				//Presenting the List First Name

				firstNames := helper.GetFirstNames()
				fmt.Printf("The first names of booking are : %v\n ", firstNames)

				//Presenting All Users
				getAllUserList := helper.GetAllUser()

				fmt.Printf("The users are  : %v\n ", getAllUserList)

				fmt.Printf("The users from the list  are  : %v\n ", helper.GetAllUserListFromBooking())

				helper.SendTicket(userTickets, firstName, lastName, email)

				//Presenting All Email User List

				getAllEmailAdd := helper.GetAllEmail()
				fmt.Printf("The email list  are  : %v\n ", getAllEmailAdd)

				//Hahs256 all emailAddress

				helper.HashEmailEncerpt()

				// Create a MySQL database connection string

				dbUser := "root"
				dbPass := "UziNarkis5!"
				dbName := "GoUserList"
				dbHost := "localhost"
				dbPort := 3306

				dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

				createTableQuery := `
				CREATE TABLE IF NOT EXISTS users (
					id INT AUTO_INCREMENT PRIMARY KEY,
					firstName VARCHAR(50),
					lastName VARCHAR(50),
					email VARCHAR(100),
					userTickets INT

				)
			`

				// Prepare the SQL query
				insertQuery := ` INSERT INTO users (firstName, lastName, email,userTickets) VALUES (?, ?, ?, ?)`

				dbConnectionStatus := dbConnection.DbConnectionVerification(dataSourceName, createTableQuery, insertQuery, firstName, lastName, email, userTickets)

				fmt.Printf("The Connection is ok  : %v\n ", dbConnectionStatus)

				if remainingTickets == 0 {

					//end program

					fmt.Println("Our conference is booked out .Come back next year.")
					break
				}

			}

		} else {

			if !isValidName {
				fmt.Printf("First ane or last Name is short , please fill out again ")

			}
			if !isValidEmail {
				fmt.Printf("The email you entered isnt correct  ")

			}

			if !isValidTicketNumber {
				fmt.Printf("The amount of ticet is incorrect  ")

			}

			fmt.Printf("Your input data is invalid , try again ")
			//adding the option to continur the programing
		}

		// Define routes
		//http.HandleFunc("/UsersData", getUsersData)

		// Start server
		//log.Fatal(http.ListenAndServe(":8090", nil))

		//greetUsers()

	}
}
