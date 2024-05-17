package common

import (
	"errors"
	"fmt"
	"go-contact-manager/modules/contacts"
	"log"
	"os"
	"regexp"
)

var logger = log.New(os.Stdout, "controller", log.LstdFlags)

func ValidateData(contact contacts.ContactInfo) error {
	if ValidateName(contact.FirstName, contact.LastName) {
		if ValidateEmail(contact.Email) {
			return nil
		} else {
			return errors.New("please enter the right email")
		}
	} else {
		return errors.New("please enter the right name")
	}

}

func ValidateName(firstName string, lastName string) bool {
	regex := "^[A-Za-z]+(?: [A-Za-z]+)*$"
	re := regexp.MustCompile(regex)
	return re.MatchString(fmt.Sprintf("%s %s", firstName, lastName))
}
func ValidateEmail(email string) bool {
	regex := "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}
