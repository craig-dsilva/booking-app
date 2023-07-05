package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint =  50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var name string
	var userTickets int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&name)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v for booking %v tickets.\n", name, userTickets)
}