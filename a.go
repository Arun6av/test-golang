package main

import (
	"fmt"
	"strings"
)

func main() {
	var ConferenceName = "Go conference"

	const ConferenceTickets = 50

	var RemainingTickets = 50

	Bookings := []string{}

	fmt.Printf("ConferenceName is %T, ConferenceTickets is %T, RemainingTickets is %T\n", ConferenceName, ConferenceTickets, RemainingTickets)

	fmt.Printf("Welcome to %v booking application\n", ConferenceName)

	fmt.Printf("We have total of %v tickets and %v tickets are still available\n", ConferenceTickets, RemainingTickets)

	fmt.Println("Get your tickets to attend", ConferenceName)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		fmt.Println("Enter your First Name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your Last Name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your Email Address:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		RemainingTickets = RemainingTickets - userTickets

		Bookings = append(Bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets.You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

		fmt.Printf("%v tickets remaining for %v\n", RemainingTickets, ConferenceName)

		FirstNames := []string{}
		for _, booking := range Bookings {
			var names = strings.Fields(booking)
			FirstNames = append(FirstNames, names[0])
		}
		fmt.Printf("The firstname of the bookings are : %v\n", FirstNames)

		if RemainingTickets == 0 {
			fmt.Println("Our conference is booked out.")
			break
		}

	}

}
