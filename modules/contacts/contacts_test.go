package contacts

import (
	"errors"
	"testing"
)

func TestAddContact(t *testing.T) {
	tests := []struct {
		name        string
		contact     ContactInfo
		contactList []ContactInfo
		mockSave    func(contactList []ContactInfo) error
		wantError   bool
	}{
		{
			name:        "Success",
			contact:     ContactInfo{FirstName: "John", LastName: "Jones", Email: "john@jones.com"},
			contactList: []ContactInfo{},
			mockSave: func(contactList []ContactInfo) error {
				t.Log("Mock save successful")
				return nil
			},
			wantError: false,
		},
		{
			name:        "Error",
			contact:     ContactInfo{FirstName: "John", LastName: "Jones", Email: "john@jones.com"},
			contactList: []ContactInfo{},
			mockSave: func(contactList []ContactInfo) error {
				t.Log("Mock save error")
				return errors.New("mock error")
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saveContacts = tt.mockSave
			err := AddContact(tt.contact, tt.contactList)
			if (err != nil) != tt.wantError {
				t.Errorf("AddContact() error = %v, wantError = %v", err, tt.wantError)
			}
		})
	}
}

func TestViewContacts(t *testing.T) {
	tests := []struct {
		name        string
		contactList []ContactInfo
	}{
		{
			name:        "Testing view",
			contactList: []ContactInfo{{FirstName: "John", LastName: "Jones", Email: "John@jones.com"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("Viewing John Details")
			ViewContacts(tt.contactList)
		})
	}
}
