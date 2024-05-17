package main

import (
	"fmt"
	"go-contact-manager/modules/common"
	contacts "go-contact-manager/modules/contacts"
	"log"
	"os"
)

var contactList []contacts.ContactInfo

func main() {
	logger := log.New(os.Stdout, "controller", log.LstdFlags)

	logger.Println("Welcome to the contacts manager")
loop:
	for {
		if err := contacts.LoadData(&contactList); err != nil {
			logger.Fatalln("Error loading data")
		}
		var choice int
		logger.Println("Enter your choice")
		logger.Println("1. Add contact")
		logger.Println("2. View contacts")
		logger.Println("3. Search contact")
		logger.Println("4. Delete contact")
		logger.Println("5. Update contact")
		logger.Println("0. Exit application")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			var firstName, lastName, email string
			logger.Println("Enter the first name of the contact you want to add")
			fmt.Scanln(&firstName)
			logger.Println("Enter the last name of the contact you want to add")
			fmt.Scanln(&lastName)
			logger.Println("Enter the email of the contact you want to add")
			fmt.Scanln(&email)
			contact := contacts.ContactInfo{
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
			}
			if err := common.ValidateData(contact); err != nil {
				logger.Fatalln("Error validating data. ", err)
				break
			}
			if err := contacts.AddContact(contact, contactList); err != nil {
				logger.Fatalln("Error adding contact ", err)
			}
		case 2:
			contacts.ViewContacts(contactList)
		case 3:
			var firstName string
			logger.Println("Enter the firstname of the contact you want to search")
			fmt.Scanln(&firstName)
			if index, err := contacts.SearchContacts(firstName, contactList); err == nil {
				logger.Printf("Entry:\nFirstName: %s, LastName: %s, Email: %s\n", contactList[index].FirstName, contactList[index].LastName, contactList[index].Email)
			} else {
				logger.Fatalln(err)
			}
		case 4:
			var firstName string
			logger.Println("Enter the firstname of the contact you want to delete")
			fmt.Scanln(&firstName)
			if err := contacts.DeleteContact(firstName, contactList); err != nil {
				logger.Fatalln(err)
			}
		case 5:
			var firstName string
			logger.Println("Enter the firstname of the contact you want to update")
			fmt.Scanln(&firstName)
			if err := contacts.UpdateContact(firstName, contactList); err != nil {
				logger.Fatalln(err)
			}
		case 0:
			logger.Println("Exiting application")
			break loop
		default:
			logger.Println("Wrong Entry try again")
		}
	}
}
