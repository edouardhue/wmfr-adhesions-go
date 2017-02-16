package memberships

import (
	"testing"
	"github.com/wikimedia-france/wmfr-adhesions/civicrm"
	"github.com/wikimedia-france/wmfr-adhesions/iraiser"
	"github.com/stretchr/testify/assert"
	"time"
	"github.com/wikimedia-france/wmfr-adhesions/internal"
	"math/rand"
)

func TestGetContactWithCrmError(t *testing.T) {
	contactGetter = func(*civicrm.GetContactQuery) (*civicrm.GetContactResponse, error) {
		response := &civicrm.GetContactResponse{}
		response.IsError = 1
		return response, assert.AnError
	}
	donation := iraiser.Donation{}
	matches, err := getContact(&donation)
	assert.Equal(t, 0, matches, "No contact")
	assert.Error(t, err, "Result")
}

func TestGetContactNotFound(t *testing.T) {
	contactGetter = func(*civicrm.GetContactQuery) (*civicrm.GetContactResponse, error) {
		response := &civicrm.GetContactResponse{}
		return response, nil
	}
	donation := iraiser.Donation{}
	matches, err := getContact(&donation)
	assert.Equal(t, 0, matches, "No contact")
	assert.Error(t, err, "Result")
}

func TestGetContactSeveralResponses(t *testing.T) {
	contactGetter = func(*civicrm.GetContactQuery) (*civicrm.GetContactResponse, error) {
		response := &civicrm.GetContactResponse{}
		response.Count = 2
		return response, nil
	}
	donation := iraiser.Donation{}
	matches, err := getContact(&donation)
	assert.Equal(t, 0, matches, "No contact")
	assert.Error(t, err, "Result")
}

func TestGetContactExactlyOne(t *testing.T) {
	contactId  := rand.Int()
	contactGetter = func(*civicrm.GetContactQuery) (*civicrm.GetContactResponse, error) {
		response := &civicrm.GetContactResponse{}
		response.Count = 1
		response.Id = contactId
		return response, nil
	}
	donation := iraiser.Donation{}
	matches, err := getContact(&donation)
	assert.Equal(t, contactId, matches, "No contact")
	assert.NoError(t, err, "Result")
}

func TestCreateContact(t *testing.T) {
	contactId  := rand.Int()
	validationDate := time.Now()
	internal.Config.ContactTypeName = "human"
	internal.Config.ContactSourceName = "iraiser"
	internal.Config.LocationTypeId = rand.Int()
	donation := iraiser.Donation{
		Donator: iraiser.Donator{
			Mail: "test@example.org",
			FirstName: "First",
			LastName: "Last",
			Pseudo: "Nick",
			StreetAddress: "Address",
			City: "City",
			PostalCode: "12345",
			Country: "CY",
		},
		Campaign: iraiser.Campaign{
			AffectationCode: "Aff",
			OriginCode: "Ori",
		},
		Payment: iraiser.Payment{
			Mode: "card",
			GatewayId: "6789",
		},
		Amount: rand.Int(),
		Currency: "CUR",
		Reference: "abcd",
		ValidationDate: validationDate,
	}
	contactCreator = func(contact *civicrm.Contact) (*civicrm.CreateContactResponse, error) {
		assert.Equal(t, donation.Mail, contact.Mail, "Mail is propagated")
		assert.Equal(t, donation.FirstName, contact.FirstName, "First name is propagated")
		assert.Equal(t, donation.LastName, contact.LastName, "Last name is propagated")
		assert.Equal(t, donation.Pseudo, contact.Pseudo, "Pseudo is propagated")
		assert.Equal(t, internal.Config.ContactTypeName, contact.ContactType, "Contact type is set")
		assert.Equal(t, internal.Config.ContactSourceName, contact.Source, "Source is set")
		response := &civicrm.CreateContactResponse{}
		response.Id = contactId
		return response, nil
	}
	addressCreator = func(address *civicrm.Address) (*civicrm.CreateAddressResponse, error) {
		assert.Equal(t, address.ContactId, contactId, "Contact id is propagated")
		assert.Equal(t, address.City, donation.City, "City is propagated")
		assert.Equal(t, address.Country, donation.Country, "Country is propagated")
		assert.Equal(t, address.PostalCode, donation.PostalCode, "Postal code is propagated")
		assert.Equal(t, address.StreetAddress, donation.StreetAddress, "Address is propagated")
		assert.Equal(t, address.LocationTypeId, internal.Config.LocationTypeId, "Location type id is set")
		assert.Equal(t, address.StreetParsing, streetParsing, "Street parsing is set")
		return &civicrm.CreateAddressResponse{}, nil
	}
	id, err := createContact(&donation)
	assert.NoError(t, err, "Result")
	assert.Equal(t, contactId, id, "Contact id is returned")
}

func TestCreateContactWithCrmContactError(t *testing.T) {
	donation := iraiser.Donation{}
	contactCreator = func(contact *civicrm.Contact) (*civicrm.CreateContactResponse, error) {
		return &civicrm.CreateContactResponse{}, assert.AnError
	}
	addressCreator = func(*civicrm.Address) (*civicrm.CreateAddressResponse, error) {
		t.Fatal("Should not attempt to create an address now.")
		return &civicrm.CreateAddressResponse{}, nil
	}
	id, err := createContact(&donation)
	assert.Error(t, err, "Error")
	assert.Equal(t, 0, id, "No contact")
}

func TestCreateContactWithCrmAddressError(t *testing.T) {
	donation := iraiser.Donation{}
	contactCreator = func(contact *civicrm.Contact) (*civicrm.CreateContactResponse, error) {
		return &civicrm.CreateContactResponse{}, nil
	}
	addressCreator = func(*civicrm.Address) (*civicrm.CreateAddressResponse, error) {
		return &civicrm.CreateAddressResponse{}, assert.AnError
	}
	id, err := createContact(&donation)
	assert.Error(t, err, "Error")
	assert.Equal(t, 0, id, "No contact")
}