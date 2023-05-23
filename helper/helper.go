package helper

import "strings"

var myVar = "somevalue"

var dbUser string
var doPass string
var dbHost string
var dbName string
var dbPort string
var dataSourceName string
var dbValidConnection bool

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber

}
