// create package main
package main

//import fmt for print
import "fmt"

func main() {
	var confName string = "Go Conference"
	const confTicketNumber int = 100
	var remainingTicket int = 100
	fmt.Println("Hello, Users!")
	fmt.Printf("Welcome to %s\n", confName)
	fmt.Printf("We have %d tickets remaining\n", remainingTicket)

	var userFirstName string
	var userLstName string
	var userEmail string
	var numberOfTickets uint8

	// ask user for their name
	fmt.Println("Please enter first name")
	fmt.Scan(&userFirstName)
	fmt.Println("Please enter last name")
	fmt.Scan(&userLstName)
	fmt.Println("Please enter email")
	fmt.Scan(&userEmail)
	fmt.Println("How many tickets would you like to purchase?")
	fmt.Scan(&numberOfTickets)

}
