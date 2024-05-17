package contacts

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type ContactInfo struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
}

func LoadData(contacts *[]ContactInfo) error {
	file, err := os.Open("contacts-db.json")
	if err != nil {
		fmt.Println("Error opening file")
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
	if err := SaveContacts(contactList); err != nil {
		return err
	}
	return nil
}

// view
func ViewContacts(contactList []ContactInfo) {
	for _, contact := range contactList {
		fmt.Printf("FirstName: %s, LastName: %s, Email: %s\n", contact.FirstName, contact.LastName, contact.Email)
	}
}

// search
func SearchContacts(firstName string, contacts []ContactInfo) (int, error) {
	for index, contact := range contacts {
		if contact.FirstName == firstName {
			fmt.Println("Contact found at position", index)
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
	fmt.Println("Enter the first name of the contact you want to update")
	fmt.Scanln(&updatedFirstName)
	fmt.Println("Enter the last name of the contact you want to update")
	fmt.Scanln(&updatedLastName)
	fmt.Println("Enter the email of the contact you want to update")
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
