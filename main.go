package main

import (
	"fmt"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint =  50
	bookings := []string{}
	
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var name string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&name)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(name) >= 2
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, name)

			fmt.Printf("Thank you %v for booking %v tickets.\n", name, userTickets)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			fmt.Printf("These are all our bookings: %v\n", bookings)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Name you entered is too short")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	}
}