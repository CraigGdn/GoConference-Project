package main

import "fmt"

func greetusers() {
	fmt.Printf("Hello and welcome to the %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and we have %v remaining tickets available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}