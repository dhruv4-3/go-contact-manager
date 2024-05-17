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

func TestSearchContacts(t *testing.T) {
	tests := []struct {
		name      string
		firstName string
		contacts  []ContactInfo
		expected  error
	}{
		{
			name:      "John_Success",
			firstName: "John",
			contacts:  []ContactInfo{{FirstName: "John", LastName: "Jones", Email: "john@jones.com"}},
			expected:  nil,
		},
		{
			name:      "John_Error",
			firstName: "Sandra",
			contacts:  []ContactInfo{{FirstName: "John", LastName: "Jones", Email: "john@jones.com"}},
			expected:  errors.New("could not find the contact"),
		},
	}
	t.Log("Running tests for Searching John")
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := SearchContacts(test.firstName, test.contacts)
			if err != nil {
				t.Logf("Expected %v, Got %v", test.expected, err)
			} else {
				t.Logf("Expected %v, Got %v", test.expected, err)
			}
		})
	}
}
