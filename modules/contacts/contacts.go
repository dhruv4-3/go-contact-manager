package contacts

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

var saveContacts = SaveContacts

type ContactInfo struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
}

var logger = log.New(os.Stdout, "controller", log.LstdFlags)

func LoadData(contacts *[]ContactInfo) error {
	file, err := os.Open("contacts-db.json")
	if err != nil {
		logger.Println("Error opening file")
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(contacts); err != nil {
		return err
	}
	return nil
}

func SaveContacts(contacts []ContactInfo) error {
	file, err := os.OpenFile("contacts-db.json", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return errors.New("error opening file")
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(contacts); err != nil {
		return errors.New("error writing file contents")
	}
	return nil
}

// add
func AddContact(contact ContactInfo, contactList []ContactInfo) error {
	contactList = append(contactList, contact)
	if err := saveContacts(contactList); err != nil {
		return err
	}
	return nil
}

// view
func ViewContacts(contactList []ContactInfo) {
	for _, contact := range contactList {
		logger.Printf("FirstName: %s, LastName: %s, Email: %s\n", contact.FirstName, contact.LastName, contact.Email)
	}
}

// search
func SearchContacts(firstName string, contacts []ContactInfo) (int, error) {
	for index, contact := range contacts {
		if contact.FirstName == firstName {
			logger.Println("Contact found at position", index)
			return index, nil
		}
	}
	return 0, errors.New("could not find the contact")
}

// delete
func DeleteContact(firstName string, contacts []ContactInfo) error {
	index, err := SearchContacts(firstName, contacts)
	if err != nil {
		return err
	}
	contacts = append(contacts[:index], contacts[index+1:]...)
	if err := SaveContacts(contacts); err != nil {
		return err
	}
	return nil
}

// update
func UpdateContact(firstName string, contacts []ContactInfo) error {
	err := DeleteContact(firstName, contacts)
	if err != nil {
		return err
	}
	var updatedFirstName, updatedLastName, updatedEmail string
	logger.Println("Enter the first name of the contact you want to update")
	fmt.Scanln(&updatedFirstName)
	logger.Println("Enter the last name of the contact you want to update")
	fmt.Scanln(&updatedLastName)
	logger.Println("Enter the email of the contact you want to update")
	fmt.Scanln(&updatedEmail)

	var contact = ContactInfo{
		FirstName: updatedFirstName,
		LastName:  updatedLastName,
		Email:     updatedEmail,
	}

	if err := LoadData(&contacts); err != nil {
		return err
	}
	if err := AddContact(contact, contacts); err != nil {
		return err
	}
	return nil
}
