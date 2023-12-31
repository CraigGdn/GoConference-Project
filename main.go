package main

import (
	"fmt"
<<<<<<< Updated upstream
	"goConference/helper"
	"sync"
=======
	"strings"

//	"golang.org/x/tools/go/analysis/passes/bools"
)

func main() { //main is entry point to the program

	greetusers()

	//fmt.Printf("This section will outline what the data types are:\n\n - conferenceName is '%T'\n - conferenceTickets is '%T'\n - remainingTickets is '%T'\n\n", conferenceName, conferenceTickets, remainingTickets)

	//fmt.Printf("Hello and welcome to the %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and we have %v remaining tickets available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	//for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets) //using 'V' after helper. will import the user package 

		if isValidName && isValidEmail && isValidTicketNumber {

<<<<<<< Updated upstream
			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1) // sets number of goroutines to wait for (increases the counter by the provided number)
			go sendTicket(userTickets, firstName, lastName, email) // placing 'go' at the start of this line makes the program become concurrent as it starts a ne go routine
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

=======
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v Tickets are remaining for the %v\n", remainingTickets, conferenceName)
			/*
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are %v\n \n", firstNames)
			*/
			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year")
				//break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("The email you entered is incorrect")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
			fmt.Printf("Your input data is invalid, try again\n")
		}
		wg.Wait()
	//}
}
