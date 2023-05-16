package main

import (
	"GOLNGCOURSE/helper"
	"encoding/json"
	"fmt"

	//"log"
	"net/http"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets int = 50

type getAllUserList []getAllUserList
type getAllEmailAdd []getAllEmailAdd

var valideUserExsitingResult bool
var isValidUserExsiting bool = true
var ifValideUserExsiting string
var amountTicket uint
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName      string
	lastName       string
	email          string
	numberOfTicket uint
}

var UsersData []UserData

type UserEmail []UserEmail

func main() {

	fmt.Printf("We have total of %v tickets and %v\n are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend")

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			//Validation of User Exsiting

			valideUserExsitingResult := valideUserExsiting(firstName)
			fmt.Printf("Indication if user is exsiting   : %v\n ", valideUserExsitingResult)

			//Presenting the Ticket Amount

			amountTicket = bookTicket(userTickets, firstName, lastName, email)

			fmt.Printf("The amount of ticket left  : %v\n ", amountTicket)

			//Presenting the List First Name

			firstNames := getFirstNames()
			fmt.Printf("The first names of booking are : %v\n ", firstNames)

			//Presenting All Users
			getAllUserList := getAllUser()

			fmt.Printf("The users are  : %v\n ", getAllUserList)

			fmt.Printf("The users from the list  are  : %v\n ", getAllUserListFromBooking())

			sendTicket(userTickets, firstName, lastName, email)

			//Presenting All Email User List

			getAllEmailAdd := getAllEmail()
			fmt.Printf("The email list  are  : %v\n ", getAllEmailAdd)

			if remainingTickets == 0 {

				//end program

				fmt.Println("Our conference is booked out .Come back next year.")
				break
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

func valideUserExsiting(firstName string) bool {

	listofAllName := getAllUserListFromBooking()

	fmt.Printf("Testing The list  of All Name are  : %v\n ", listofAllName)

	fmt.Printf("Testing The list  of firstName   : %v\n ", firstName)

	for _, listofAllName := range listofAllName {

		if listofAllName == firstName {

			fmt.Printf("TestingAgain The list  of firstName\n   : %v\n ", firstName)

			fmt.Printf("user is exsiting\n  ")

			fmt.Printf("TestingAgain The list  of isValidUSerExsiting \n   : %v\n ", isValidUserExsiting)

			isValidUserExsiting := true
			return isValidUserExsiting

		}

	}

	fmt.Printf("user isnt exsiting\n  ")
	isValidUserExsiting := false
	return isValidUserExsiting

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v\n are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend")
}

func getFirstNames() []string {

	firstNames := []string{}

	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)

	}

	return firstNames

}

func getUserInput() (string, string, string, uint) {

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Printf("Enter your First anme :")
		fmt.Scan(&firstName)

		fmt.Printf("Enter your Last anme :")
		fmt.Scan(&lastName)

		fmt.Printf("Enter your email add :")
		fmt.Scan(&email)

		fmt.Printf("Enter your of ticket  :")
		fmt.Scan(&userTickets)

		return firstName, lastName, email, userTickets

	}
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) uint {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{

		firstName:      firstName,
		lastName:       lastName,
		email:          email,
		numberOfTicket: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("list of booking is %v \n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will reacive a confirantion email %v at\n", firstName, userTickets, lastName, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
	fmt.Printf("userData %v \n", bookings)

	return remainingTickets

}

func getAllUser() []UserData {

	return bookings
}

func getAllUserListFromBooking() []string {

	userNameList := []string{}

	for _, booking := range bookings {

		userNameList = append(userNameList, booking.firstName)

	}

	return userNameList

}

func getAllEmail() []string {

	UserEmail := []string{}

	for _, booking := range bookings {

		UserEmail = append(UserEmail, booking.email)

	}

	return UserEmail

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTickets, firstName, lastName)

	fmt.Printf("Sending ticket %v to email address %v\n ", ticket, email)

}

func getUsersData(w http.ResponseWriter, r *http.Request) {
	// Handle GET request for /products
	if r.Method == "GET" {
		// Encode products as JSON and write to response
		json.NewEncoder(w).Encode(UsersData)
	}
}
