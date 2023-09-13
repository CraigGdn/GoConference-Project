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

		remainingTickets = remainingTickets - userTickets
		//bookings[0] = firstName + " " + lastName //this is used for an array but not needed when using slices
		bookings = append(bookings, firstName+" "+lastName) //this is a slice

		/*
			fmt.Printf("The whole array: %v\n", bookings) //this will print the value with square brackets and will show upto 50 spaces
			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("Slice type: %T\n", bookings)        //this will print the length (50) and the type (string)
			fmt.Printf("Slice length: %v\n", len(bookings)) //this will print the length of the array (50)
		*/
		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v Tickets are remaining for the %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			//var firstName = names[0]
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of bookings are %v\n \n", firstNames)
	}
}

//Currently on if...Else Statements and Boolean 1:24:24
