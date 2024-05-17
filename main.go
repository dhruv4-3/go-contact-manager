package main

import (
	"fmt"
	"go-contact-manager/modules/common"
	contacts "go-contact-manager/modules/contacts"
)

var contactList []contacts.ContactInfo

func main() {
	fmt.Println("Welcome to the contacts manager")
loop:
	for {
		if err := contacts.LoadData(&contactList); err != nil {
			fmt.Println("Error loading data")
		}
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
			var firstName, lastName, email string
			fmt.Println("Enter the first name of the contact you want to add")
			fmt.Scanln(&firstName)
			fmt.Println("Enter the last name of the contact you want to add")
			fmt.Scanln(&lastName)
			fmt.Println("Enter the email of the contact you want to add")
			fmt.Scanln(&email)
			contact := contacts.ContactInfo{
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
			}
			if err := common.ValidateData(contact); err != nil {
				fmt.Println("Error validating data. ", err)
				break
			}
			if err := contacts.AddContact(contact, contactList); err != nil {
				fmt.Println("Error adding contact ", err)
			}
		case 2:
			contacts.ViewContacts(contactList)
		// 	// TODO Implement all missing functions
		case 3:
			var firstName string
			fmt.Println("Enter the firstname of the contact you want to search")
			fmt.Scanln(&firstName)
			if index, err := contacts.SearchContacts(firstName, contactList); err == nil {
				fmt.Printf("Entry:\nFirstName: %s, LastName: %s, Email: %s\n", contactList[index].FirstName, contactList[index].LastName, contactList[index].Email)
			} else {
				fmt.Println(err)
			}
		case 4:
			var firstName string
			fmt.Println("Enter the firstname of the contact you want to delete")
			fmt.Scanln(&firstName)
			if err := contacts.DeleteContact(firstName, contactList); err != nil {
				fmt.Println(err)
			}
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
