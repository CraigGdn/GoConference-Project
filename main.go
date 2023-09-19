package main

import (
	"fmt"
	"strings"
)

func main() { //main is entry point to the program

	conferenceName := "'Go Conference'" // the : causes Go to figure out that the line is a variable. This cannot be done for const or var where the uint is used
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{} //when using slices, the array [] should not have any numbers in the square brackets

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

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their details
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter the number tickets: ")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName) //this is a slice

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v Tickets are remaining for the %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are %v\n \n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		}

	}
}

//Currently on User Input validation 1:39:34
