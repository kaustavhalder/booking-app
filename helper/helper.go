package helper

import "fmt"

func GreetUser(confName string, confTickets int, remainingTickets int) {
	fmt.Println("Welcome to", confName, "Booking Application")
	fmt.Println("Get tickets to attend the conference")
	fmt.Println("Total Tickets", confTickets, "Still Available", remainingTickets)
}
