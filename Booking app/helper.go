package main

import "strings"

func ValidateUserInput(firstname string, lastname string, email string, usertickets uint) (bool, bool, bool) {
	isValidName := len(firstname) >= 2 && len(lastname) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := usertickets > 0 && usertickets <= remainingTickets
	// if we have to check the location of the conference isValidCity := city=="Spain" || city =="London"
	return isValidName, isValidEmail, isValidTicketNumber
}
