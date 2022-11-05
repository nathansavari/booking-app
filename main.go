package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

//package level variables

const conferenceTickets int = 50
var conferenceName = "GoConference"
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {
	greetUsers()

	for {

		firstName, lastName, email, userTickets :=  getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)

		// call function print firstName
		firstNames := getFirstNames()
		fmt.Printf("The first names are: %v\n", firstNames) 
		
		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year !")
			break
		}
	} else {
		if !isValidName {
			fmt.Println("Invalid first name")
		}
		if !isValidEmail {
			fmt.Println("Invalid email address")
		}
		if !isValidTicketNumber {
			fmt.Println("Invalid ticket number")
		}
	}
  }
}

func greetUsers() {
	fmt.Printf("Welcome to the %v\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here.")	
}

func getFirstNames() []string {
	firstNames := []string{}
		for _, booking := range bookings {
			firstNames = append(firstNames, booking["firstName"])
		}
		return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	
	fmt.Println(("Enter your first name:"))
	fmt.Scan(&firstName)

	fmt.Println(("Enter your last name:"))
	fmt.Scan(&lastName)

	fmt.Println(("Enter your email address:"))
	fmt.Scan(&email)

	fmt.Println(("Enter your number of tickets wanted:"))
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTicket( userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for a user
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)


	fmt.Printf("Thank you %v %v for booking %v tickets. Your email is %v right ?\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}