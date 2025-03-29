package main

import "fmt"

func main() {

	conferenceName := "Go conference"
	const conferenceTickets int = 50
	var remainingTickets = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T\n", conferenceName, remainingTickets)

	fmt.Printf("Welcome to   %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTicket int
	var firstName string
	var lastName string
	var email string
	//ask user for their name
	fmt.Println("Enter your name:")
	fmt.Scan(&userName)

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email:")
	fmt.Scan(&email)

	fmt.Println("Enter your name:")
	fmt.Scan(&userName)

	userTicket = 2

	fmt.Printf("USER %v BOOKED %v tickets. %v first name; %v last name	; %v email", userName, userTicket, firstName, lastName, email)
}
