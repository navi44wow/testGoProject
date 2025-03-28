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
	//ask user for their name
	fmt.Println("Enter your name:")
	fmt.Scan(&userName)

	userTicket = 2

	fmt.Printf("USER %v BOOKED %v tickets.\n", userName, userTicket)
}
