package main

import "fmt"

func main() {
	var confName = "Go Conference"
	const confTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Println("Welcome to", confName, "Booking Application")
	fmt.Println("Get tickets to attend the conference")
	fmt.Println("Total Tickets", confTickets, "Still Available", remainingTickets)

	for {
		var userName string
		var userTickets int
		// Get user input
		fmt.Println("Enter your name")
		fmt.Scan(&userName)

		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - uint(userTickets)
		bookings = append(bookings, userName)

		fmt.Printf("User %v bought %v tickets\n", userName, userTickets)
		fmt.Printf("Remaining Tickets in the conference %v\n", remainingTickets)
		fmt.Printf("Total Bookings %v\n", bookings)

	}

}
