package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"time"
)

var myVar = "somevalue"
var conferenceName = "Go Conference"

var conferenceTickets int = 50

var remainingTickets uint = 50
var dbUser string
var doPass string
var dbHost string
var dbName string
var dbPort string
var dataSourceName string
var dbValidConnection bool
var isValidUserExsiting bool = true
var bookings = make([]UserData, 0)

type UserData struct {
	FirstName      string
	LastName       string
	Email          string
	NumberOfTicket uint
}

var UsersData []UserData

type UserEmail []UserEmail

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber

}
func ValideUserExsiting(firstName string) bool {

	listofAllName := GetAllUserListFromBooking()

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

func GetAllUserListFromBooking() []string {

	userNameList := []string{}

	for _, booking := range bookings {

		userNameList = append(userNameList, booking.FirstName)

	}

	return userNameList

}

func GreetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v\n are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend")
}

func GetFirstNames() []string {

	firstNames := []string{}

	for _, booking := range bookings {

		firstNames = append(firstNames, booking.FirstName)

	}

	return firstNames

}

func HashEmailEncerpt() {

	listOfAllEmails := GetAllEmail()
	hash := HashStringSHA256(listOfAllEmails)
	fmt.Println(hash)
}

func HashStringSHA256(listOfAllEmails []string) string {

	hasher := sha256.New()
	for _, listOfAllEmails := range listOfAllEmails {
		hasher.Write([]byte(listOfAllEmails))

	}
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash

}

func GetUserInput() (string, string, string, uint) {

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Printf("Enter your First name :")
		fmt.Scan(&firstName)

		fmt.Printf("Enter your Last name :")
		fmt.Scan(&lastName)

		fmt.Printf("Enter your email add :")
		fmt.Scan(&email)

		fmt.Printf("Enter your of ticket  :")
		fmt.Scan(&userTickets)

		return firstName, lastName, email, userTickets

	}
}

func BookTicket(userTickets uint, firstName string, lastName string, email string) UserData {

	remainingTickets = remainingTickets - userTickets

	var userData = UserData{

		FirstName:      firstName,
		LastName:       lastName,
		Email:          email,
		NumberOfTicket: userTickets,
	}

	fmt.Printf("list of userData is %v \n", userData)
	bookings = append(bookings, userData)
	fmt.Printf("list of booking is %v \n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will reacive a confirantion email %v at\n", firstName, userTickets, lastName, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
	fmt.Printf("userData %v \n", bookings)

	return userData

}

func GetAllUser() []UserData {

	return bookings
}

func GetAllEmail() []string {

	UserEmail := []string{}

	for _, booking := range bookings {

		UserEmail = append(UserEmail, booking.Email)

	}

	return UserEmail

}

func SendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTickets, firstName, lastName)

	fmt.Printf("Sending ticket %v to email address %v\n ", ticket, email)

}

//func GetUsersData(w http.ResponseWriter, r *http.Request) {
// Handle GET request for /products
//	if r.Method == "GET" {
// Encode products as JSON and write to response
//		json.NewEncoder(w).Encode(UsersData)
//	}
