package main

import "fmt"

func main() { //main is entry point to the program

	conferenceName := "'Go Conference'" // the : causes Go to figure out that the line is a variable. This cannot be done for const or var where the uint is used
	const conferenceTickets = 50
	var remainingTickets uint = 50
	/*
	   This section was used with Println and is left here to show how the code looked without Printf

	   	fmt.Println("Hello and welcome to the", conferenceName, "booking application. ")
	   	fmt.Println("We have a total of", conferenceTickets, "and we have", remainingTickets, "remaining available.")
	   	fmt.Println("Get your tickets here to attend.")
	*/

	// Below can outline some data-types being used
	fmt.Printf("This section will outline what the data types are:\n\n - conferenceName is '%T'\n - conferenceTickets is '%T'\n - remainingTickets is '%T'\n\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Hello and welcome to the %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and we have %v remaining tickets available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	// ask user for their details
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
}





//Currently on Getting user input section in Youtube 45:19
