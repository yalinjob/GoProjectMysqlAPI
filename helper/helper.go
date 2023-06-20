package helper

import (
	//	"crypto/sha256"
	//	"database/sql"
	//	"encoding/hex"
	//	"fmt"
	//	"strings"
	//	"time"

	"GOLNGCOURSE/dbConfig"

	"github.com/jinzhu/gorm"
)

var myVar = "somevalue"
var conferenceName = "Go Conference"

var conferenceTickets int = 50

var remainingTickets uint = 50

var dataSourceName string
var dbValidConnection bool

//var isValidUserExsiting bool = true
//var bookings = make([]UserData, 0)

var db *gorm.DB

type UserData struct {
	gorm.Model
	FirstName string `json:"FirstName" `
	//FirstName      string `gorm: "" json:"FirstName" `
	LastName       string `json:"LastName" `
	Email          string `json:"Email" `
	NumberOfTicket uint   `json:"NumberOfTicket" `
}

//var UsersData []UserData

//type UserEmail []UserEmail

//func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
//isValidName := len(firstName) >= 2 && len(lastName) >= 2
//isValidEmail := strings.Contains(email, "@")
//	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

//	return isValidName, isValidEmail, isValidTicketNumber

//}
//func ValideUserExsiting(firstName string) bool {

//listofAllName := GetAllUserListFromBooking()

//	fmt.Printf("Testing The list  of All Name are  : %v\n ", listofAllName)

//	fmt.Printf("Testing The list  of firstName   : %v\n ", firstName)
//
//	for _, listofAllName := range listofAllName {

//		if listofAllName == firstName {

//		fmt.Printf("TestingAgain The list  of firstName\n   : %v\n ", firstName)

//		fmt.Printf("user is exsiting\n  ")

//		fmt.Printf("TestingAgain The list  of isValidUSerExsiting \n   : %v\n ", isValidUserExsiting)

//		isValidUserExsiting := true
//		return isValidUserExsiting
//
//	}

//}

//	fmt.Printf("user isnt exsiting\n  ")
//	isValidUserExsiting := false
//	return isValidUserExsiting

//}

//func GetAllUserListFromBooking() []string {

//	userNameList := []string{}

//	for _, booking := range bookings {

//	userNameList = append(userNameList, booking.FirstName)

//}

//	return userNameList

//}

//func GreetUsers() {
//	fmt.Printf("Welcome to %v booking application\n", conferenceName)
//	fmt.Printf("We have total of %v tickets and %v\n are still available.\n", conferenceTickets, remainingTickets)
//	fmt.Printf("Get your tickets here to attend")
//}

//func GetFirstNames() []string {

//	firstNames := []string{}

//	for _, booking := range bookings {

//		firstNames = append(firstNames, booking.FirstName)

//	}

//	return firstNames

//}

//func HashEmailEncerpt() {

//listOfAllEmails := GetAllEmail()
//	hash := HashStringSHA256(listOfAllEmails)
//	fmt.Println(hash)
//}

//func HashStringSHA256(listOfAllEmails []string) string {

//hasher := sha256.New()
//for _, listOfAllEmails := range listOfAllEmails {
//	hasher.Write([]byte(listOfAllEmails))

//}
//hash := hex.EncodeToString(hasher.Sum(nil))
//return hash

//}

//func GetUserInput() (string, string, string, uint) {

//for {
//	var firstName string
//	var lastName string
//	var email string
//	var userTickets uint

//	fmt.Printf("Enter your First name :")
//	fmt.Scan(&firstName)

//	fmt.Printf("Enter your Last name :")
//	fmt.Scan(&lastName)

//	fmt.Printf("Enter your email add :")
//	fmt.Scan(&email)

//	fmt.Printf("Enter your of ticket  :")
//	fmt.Scan(&userTickets)

//	return firstName, lastName, email, userTickets

//	}
//}

//func BookTicket(userTickets uint, firstName string, lastName string, email string) UserData {

//	remainingTickets = remainingTickets - userTickets

//	var userData = UserData{

//	FirstName:      firstName,
//	LastName:       lastName,
//	Email:          email,
//	NumberOfTicket: userTickets,
//}

//fmt.Printf("list of userData is %v \n", userData)
//bookings = append(bookings, userData)
//fmt.Printf("list of booking is %v \n", bookings)

//fmt.Printf("Thank you %v %v for booking %v tickets. You will reacive a confirantion email %v at\n", firstName, userTickets, lastName, email)
//fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
//fmt.Printf("userData %v \n", bookings)

//return userData

//}

//func GetAllUserFirst() []UserData {

//	return bookings
//}

//func GetAllEmail() []string {

//UserEmail := []string{}

//for _, booking := range bookings {

//	UserEmail = append(UserEmail, booking.Email)

//}

//	return UserEmail
//
//}

//func SendTicket(userTickets uint, firstName string, lastName string, email string) {
//	time.Sleep(10 * time.Second)
//	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTickets, firstName, lastName)

//fmt.Printf("Sending ticket %v to email address %v\n ", ticket, email)

//func DbConnectionVerification() {

//	dbUser := "root"
//	dbPass := "UziNarkis5!"
//	dbName := "GoUserList"
//	dbHost := "localhost"
//	dbPort := 3306
//
//	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

//	db, err := sql.Open("mysql", dataSourceName)
//	if err != nil {
//		panic(err.Error())
//	}
//	defer db.Close()
//
// Attempt to ping the database to verify the connection
//	err = db.Ping()
//	if err != nil {
//		panic(err.Error())
//	}

//	fmt.Println("Successfully connected to the MySQL database!")

//}

func init() {

	dbConfig.DbConnectionVerification()
	db = dbConfig.GetDB()
	db.AutoMigrate(&UserData{})

}

func GetUserByID(Id int64) (*UserData, *gorm.DB) {

	var getUserById UserData
	db := db.Where("id=? ", Id).Find(&getUserById)
	return &getUserById, db

}

func (ud *UserData) CreateUser() *UserData {

	db.NewRecord(ud)
	db.Create(&ud)
	return ud

}

func GetAllUser() []UserData {

	var UserDatas []UserData
	db.Find(&UserDatas)
	return UserDatas

}

func DeleteUser(ID int64) UserData {

	var userData UserData
	db.Where("ID=?", ID).Delete(userData)
	return userData

}
