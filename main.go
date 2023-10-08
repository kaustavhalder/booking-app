package main

import "fmt"

func main() {
	var confName = "Go Conference"
	const confTickets = 50
	var remainingTickets uint = 50

	fmt.Println("Welcome to", confName, "Booking Application")
	fmt.Println("Get tickets to attend the conference")
	fmt.Println("Total Tickets", confTickets, "Still Available", remainingTickets)

	var userName string
	var userTickets int
	// Get user input
	fmt.Println("Enter your name")
	fmt.Scan(&userName)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	remainingTickets = confTickets - uint(userTickets)

	fmt.Printf("User %v bought %v tickets\n", userName, userTickets)
	fmt.Printf("Remaining Tickets in the conference %v\n", remainingTickets)

}
