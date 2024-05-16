package main

// TODO Break down functions into smaller blocks
// TODO Proper validation and use Println to explore the error than forcing a crash

import (
	"fmt"
	"pkgs/contact-manager"
)

var contacts []contact.Contact

func main() {
	fmt.Println("Welcome to the contact manager")
	// loop:
	for {
		var choice int
		fmt.Println("Choose the operation")
		fmt.Println("1. Add contact")
		fmt.Println("2. View contacts")
		fmt.Println("3. Update contact")
		fmt.Println("4. Delete contact")
		fmt.Println("5. Search contact")

	}
}
