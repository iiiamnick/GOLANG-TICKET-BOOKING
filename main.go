package main

import "fmt"

func main() {
	var conferenceName = "GO CONFERENCE"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string
	fmt.Printf("WELCOME TO %v BOOKING SITE\n", conferenceName)
	fmt.Printf("TOTAL TICKETS ARE %v AND STILL REMAINING ARE %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Book Your Tickets Here")

	var firstName string
	var email string
	var bookedtickets uint

	fmt.Println("Enter your name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets to be booked")
	fmt.Scan(&bookedtickets)

	remainingTickets = remainingTickets - bookedtickets
	bookings = append(bookings, firstName)

	fmt.Printf("Thankyou %v for booking %v tickets. Your passes will be delivered to your email %v \n", firstName, bookedtickets, email)

	fmt.Printf("Reamaining tickets for the conference %v\n", remainingTickets)
	fmt.Printf("These are all bookings %v\n", bookings)

}
