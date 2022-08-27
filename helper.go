package main

import "strings"

// Capitalize first letter to make any function or variable public to all packages
func validateUserInput(firstName string, lastName string, email string, userTickets int) (bool, bool, bool) /* output variable types */ {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isVaildTicketNumber := userTickets <= remainingTickets && userTickets > 0

	return isValidName, isValidEmail, isVaildTicketNumber
}