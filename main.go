package main

import (
	"fmt"
)

func main() {

	var ConferenceName = "Go Conference"
	const ConferenceTickets = 50
	var RemTickets uint = 50
	var bookings []string

	fmt.Printf("ConferenceName is %T, ConferenceTickets is %T, RemTickets is %T \n", ConferenceName, ConferenceTickets, RemTickets)
	fmt.Printf("Welcome to %v booking page!\n", ConferenceName)
	fmt.Println("Get your tickets here to attend")

	var firstname string
	var lastname string
	var email string
	var usertickets uint
	//ask user their name
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstname)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastname)

	fmt.Print("Enter your Email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&usertickets)

	RemTickets = RemTickets - usertickets
	// bookings[0] = firstname + " " + lastname
	bookings = append(bookings, firstname+" "+lastname)

	fmt.Printf("The Whole Slice: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Slice type: %T\n", bookings)
	fmt.Printf("Slice Length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will get your confirmation email on %v\n", firstname, lastname, usertickets, email)
	fmt.Printf("%v tickets are remaining for %v\n ", RemTickets, ConferenceName)

	fmt.Printf("these are all our bookings: %v\n", bookings)
}
