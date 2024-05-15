package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func addContact(firstName string, lastName string, email string) error {
	fmt.Println("Adding new contact")
	var newContact = Contact{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
	data, err := json.Marshal(newContact)
	if err != nil {
		panic("Unable to marshall contact")
	}
	file, err := os.OpenFile("contact-db.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic("Error writing file")
	}
	if _, err := file.Write(data); err != nil {
		panic("Could not write into file")
	}
	if err := file.Close(); err != nil {
		panic("Could not close file")
	}
	return nil
}

func viewContacts() {
	fr, err := os.ReadFile("contact-db.txt")
	if err != nil {
		panic("Error reading file")
	}
	var contacts []Contact
	if err := json.Unmarshal(fr, &contacts); err != nil {
		panic("Could not unmarshall data from file")
	}
	for _, contact := range contacts {
		fmt.Printf("FirstName: %s\t LastName: %s\t Email: %s\n", contact.FirstName, contact.LastName, contact.Email)
	}
}

func searchContact(firstName string) error {
	var contacts []Contact
	fr, err := os.ReadFile("contact-db.txt")
	if err != nil {
		panic("Unable to read file")
	}
	if err := json.Unmarshal(fr, &contacts); err != nil {
		panic("Unable to unmarshall file contents")
	}
	for _, contact := range contacts {
		if contact.FirstName == firstName {
			fmt.Println("Contact found...")
			fmt.Printf("FirstName: %s\nLastName: %s\nEmail: %s", contact.FirstName, contact.LastName, contact.Email)
		}
	}
	return nil
}

func deleteContact(firstName string) error {
	var contacts []Contact
	var newContacts []Contact
	fr, err := os.ReadFile("contact-db.txt")
	if err != nil {
		panic("Could not open file")
	}
	if err := json.Unmarshal(fr, &contacts); err != nil {
		panic("Could not unmarshall file contents")
	}
	for _, contact := range contacts {
		if contact.FirstName == firstName {
			fmt.Println("Contact found in database. Proceeding to delete")
			continue
		} else {
			newContacts = append(newContacts, contact)
		}
	}

	data, err := json.Marshal(newContacts)
	if err != nil {
		panic("Could not marshall new contacts")
	}

	if err := os.WriteFile("contact-db.txt", data, 0644); err != nil {
		panic("Could not write file after delete")
	}
	return nil
}

func main() {
loop:
	for {
		fmt.Println("Welcome to the contact manager")
		fmt.Println("Enter your choice")
		fmt.Println("1. Add new contact")
		fmt.Println("2. View contacts")
		fmt.Println("3. Search for contact")
		fmt.Println("4. Delete a contact")
		fmt.Println("0. Exit application")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			var firstName string
			var lastName string
			var email string
			fmt.Println("Enter firstName of the contact")
			fmt.Scanln(&firstName)
			fmt.Println("Enter lastName of the contact")
			fmt.Scanln(&lastName)
			fmt.Println("Enter email of the contact")
			fmt.Scanln(&email)
			if len(firstName) < 1 || len(lastName) < 1 || len(email) < 1 {
				panic("Please enter the right values")
			}
			if err := addContact(firstName, lastName, email); err != nil {
				panic("Could not add contact")
			}
		case 2:
			fmt.Println("Viewing Contacts")
			viewContacts()
		case 3:
			fmt.Println("Enter the first name of the contact you want to search for")
			var firstName string
			fmt.Scanln(&firstName)
			if err := searchContact(firstName); err != nil {
				panic("Error searching contact")
			}
		case 4:
			fmt.Println("Enter the first name of the contact you want to delete")
			var firstName string
			fmt.Scanln(&firstName)
			if err := deleteContact(firstName); err != nil {
				panic("Could not delete contact")
			}
		case 0:
			break loop
		default:
			panic("Invalid choice")
		}
	}
}
