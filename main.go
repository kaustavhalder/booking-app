package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

const confTickets = 50

var confName = "Go Conference"
var remainingTickets uint = 50
var bookings []string

func main() {
	helper.GreetUser(confName, confTickets, int(remainingTickets))

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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isUserTickets := userTickets > 0 && userTickets <= 50

		if isValidName && isUserTickets {
			remainingTickets = remainingTickets - uint(userTickets)
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("User %v bought %v tickets\n", firstName+" "+lastName, userTickets)
			fmt.Printf("Remaining Tickets in the conference %v\n", remainingTickets)

			printFirstName(bookings)
			if remainingTickets == 0 {
				fmt.Println("Tickets have been sold out")
				break
			}

		} else {
			// fmt.Printf("we only have %v tickets available, you cant book %v tickets\n", remainingTickets, userTickets)
			if !isValidName {
				fmt.Println("you have entered wrong first name or last name")
			}
			if !isUserTickets {
				fmt.Println("no of tickets have been entered wrongly")
			}
		}
	}
}


func printFirstName(bookings []string) {
	firstNames := []string{}
	for _, val := range bookings {
		var name = strings.Fields(val)
		firstNames = append(firstNames, name[0])
	}
	fmt.Printf("Bookings made by %v\n", firstNames)
}
