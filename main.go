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

	fmt.Printf("This section will outline what the data types are:\n\n - conferenceName is '%T'\n - conferenceTickets is '%T'\n - remainingTickets is '%T'\n\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Hello and welcome to the %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and we have %v remaining tickets available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

//Currently on Getting user input section in Youtube 45:19
