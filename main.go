package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference" // package level variables, can be used anywhere
const conferenceTickets = 50
var remainingTickets = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	userTickets int
}

var wg sync.WaitGroup // Waitgroup to not let the main program exit without completing all the go routines (threads)

func main(){
	greetUsers()
	
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isVaildTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

	if isValidName && isValidEmail && isVaildTicketNumber {
		
		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email) /* go keyword creates a new thread for the function to be executed
																	to not stop the flow of the main program */
		printFirstNames()

		if remainingTickets <= 0 {
			fmt.Println("All the tickets for now are booked, Come back next year :D")
		}
	} else {
		if !isValidName {
			fmt.Println("Name Entered by you is too short")
		}

		if !isValidEmail {
			fmt.Println("Email you entered does not have @ sign")
		}

		if !isVaildTicketNumber {
			fmt.Println("Number of tickets entered are invalid")
		}
		fmt.Println("Please Try Again.")
	}
	wg.Wait()
}

func greetUsers(){
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames() {
	fmt.Println("Your first name of bookings are:- ")

	for _, booking := range bookings {
		fmt.Println(booking.firstName)
	}
}

func getUserInput() (string, string, string, int){
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName) 

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining of %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#####################")
	fmt.Printf("Sending ticket:\n%v\nto email address: %v\n", ticket, email)
	fmt.Println("#####################")
	wg.Done()
}