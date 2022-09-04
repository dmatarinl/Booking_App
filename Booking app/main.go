package main

import (
	"fmt"
	"time"
)

//package level variables

var confName = "Go conference"

var remainingTickets uint = 50

const confTickets int = 50

var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	NumberofTickets uint
}

func main() {

	greetusers()

	for remainingTickets > 0 && len(bookings) < 50 {

		firstname, lastname, email, usertickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstname, lastname, email, usertickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(usertickets, firstname, lastname, email)

			//go in front for concurrency, multiple threads or orders at the same time
			go sendTicket(usertickets, firstname, lastname, email)

			firstnames := getfirstnames()
			fmt.Printf("The first name of bookings are: %v\n", firstnames)

			if remainingTickets == 0 {
				//end the program
				fmt.Println("Our conference is booked out. Come back next year :)")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered does not contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}

		}

	}

}

func greetusers() {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here")
}

func getfirstnames() []string {
	firstnames := []string{}
	// _ underscore -> ignore the variable we don't  want to use (index)
	for _, booking := range bookings {

		firstnames = append(firstnames, booking.firstName)

	}
	return firstnames

}

func getUserInput() (string, string, string, uint) {
	var firstname string
	var lastname string
	var email string
	var usertickets uint

	//asking user for personal data

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstname)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastname)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&usertickets)

	return firstname, lastname, email, usertickets
}

func bookTicket(usertickets uint, firstname string, lastname string, email string) {
	remainingTickets = remainingTickets - usertickets

	//create a map for a user

	var userData = UserData{
		firstName:       firstname,
		lastName:        lastname,
		email:           email,
		NumberofTickets: usertickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstname, lastname, usertickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confName)
}

func sendTicket(usertickets uint, firstName string, lastName string, email string) {
	time.Sleep(35 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", usertickets, firstName, lastName)
	fmt.Println("$$$$$$$$$$$$$$$$$")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("$$$$$$$$$$$$$$$$$")
}
