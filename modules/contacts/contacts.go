package contacts

//TODO Have to learn to write json file
import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
)

type ContactInfo struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
}

func AddContact() error {
	var firstName, lastName, email string
	fmt.Println("Enter the first name of the contact")
	fmt.Scanln(&firstName)
	fmt.Println("Enter the last name of the contact")
	fmt.Scanln(&lastName)
	fmt.Println("Enter the email of the contact")
	fmt.Scanln(&email)

	var contacts []ContactInfo

	var contact = ContactInfo{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	if err := validateData(contact); err != nil {
		fmt.Println("Error validating data")
		return err
	} else {
		contacts = append(contacts, contact)
		saveContact(contacts)
		return nil
	}
}

func validateData(contact ContactInfo) error {
	if validateName(contact.FirstName, contact.LastName) {
		if validateEmail(contact.Email) {
			return nil
		} else {
			return errors.New("please enter the right email")
		}
	} else {
		return errors.New("please enter the right name")
	}

}

func validateName(firstName string, lastName string) bool {
	regex := "^[A-Za-z]+(?: [A-Za-z]+)*$"
	re := regexp.MustCompile(regex)
	return re.MatchString(fmt.Sprintf("%s %s", firstName, lastName))
}
func validateEmail(email string) bool {
	regex := "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

func saveContact(contacts []ContactInfo) {
	file, err := os.OpenFile("contacts-db.json", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(contacts); err != nil {
		fmt.Println("Cannot write json to file")
	}
}

func ViewContact() error {
	var firstName string
	var contacts []ContactInfo
	fmt.Println("Enter the first name of the contact you want to search")
	fmt.Scanln(&firstName)
	loadData(&contacts)
	for _, contact := range contacts {
		if contact.FirstName == firstName {
			fmt.Println("Contact found")
			fmt.Printf("First Name: %s\t Last Name: %s\t Email: %s\n", contact.FirstName, contact.LastName, contact.Email)
			return nil
		}
	}
	return errors.New("could not find the contact")
}

func loadData(contacts *[]ContactInfo) {
	file, err := os.OpenFile("contacts-db.json", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(contacts); err != nil && err != io.EOF {
		fmt.Println("Error unmarshalling data from file")
	}
}
