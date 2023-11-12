package main

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
	firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}