package main

import (
	"fmt"
	"strings"
)

func main() {
	var confName = "Go Conference"
	const confTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Println("Welcome to", confName, "Booking Application")
	fmt.Println("Get tickets to attend the conference")
	fmt.Println("Total Tickets", confTickets, "Still Available", remainingTickets)

	for {
		var firstName string
		var lastName string
		var userTickets int
		// Get user input
		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)

		// Perform operations
		remainingTickets = remainingTickets - uint(userTickets)
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("User %v bought %v tickets\n", firstName+" "+lastName, userTickets)
		fmt.Printf("Remaining Tickets in the conference %v\n", remainingTickets)

		firstNames := []string{}
		for _, val := range bookings{
			var name = strings.Fields(val)
			firstNames = append(firstNames, name[0])
		}
		fmt.Printf("Bookings made by %v\n", firstNames)

	}

}
