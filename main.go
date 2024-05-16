package main

import (
	"fmt"
	contacts "go-contact-manager/modules/contacts"
)

var contactList []contacts.ContactInfo

func main() {
	fmt.Println("Welcome to the contacts manager")
loop:
	for {
		var choice int
		fmt.Println("Enter your choice")
		fmt.Println("1. Add contact")
		fmt.Println("2. View contacts")
		fmt.Println("3. Search contact")
		fmt.Println("4. Delete contact")
		fmt.Println("5. Update contact")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			if err := contacts.AddContact(); err != nil {
				fmt.Println("Error adding contact")
			}
		case 2:
			if err := contacts.ViewContact(); err != nil {
				fmt.Println("Error viewing contact")
			}
			// TODO Implement all missing functions
		// case 3:
		// 	if err := contacts.SearchContact(); err != nil {
		// 		fmt.Println("Error searching contact")
		// 	}
		// case 4:
		// 	if err := contacts.DeleteContact(); err != nil {
		// 		fmt.Println("Error deleting contact")
		// 	}
		// case 5:
		// 	if err := contacts.UpdateContact(); err != nil {
		// 		fmt.Println("Error updating contact")
		// 	}
		case 0:
			fmt.Println("Exiting application")
			break loop
		default:
			fmt.Println("Wrong Entry try again")
		}
	}
}
