package main

import "fmt"

func main() {
	var confName = "Go Conference"
	const confTickets = 50
	var remainingTickets = 50
	fmt.Println("Welcome to", confName, "Booking Application")
	fmt.Println("Get tickets to attend the conference")
	fmt.Println("Total Tickets", confTickets, "Still Available", remainingTickets)

}
