package main

import "fmt"

func main() { //main is entry point to the program

	var conferenceName = "'Go Conference'"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Hello and welcome to the", conferenceName, "booking application. ")
	fmt.Println("We have a total of", conferenceTickets, "and we have", remainingTickets, "remaining available.")
	fmt.Println("Get your tickets here to attend.")
}

//Currently on Format output section in Youtube 31:01
