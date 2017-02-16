package memberships

import (
	"github.com/wikimedia-france/wmfr-adhesions/iraiser"
	"github.com/wikimedia-france/wmfr-adhesions/civicrm"
	"github.com/wikimedia-france/wmfr-adhesions/internal"
)

var contactGetter func(*civicrm.GetContactQuery) (*civicrm.GetContactResponse, error) = civicrm.GetContact
var contactCreator func(*civicrm.Contact) (*civicrm.CreateContactResponse, error) = civicrm.CreateContact
var addressCreator func(*civicrm.Address) (*civicrm.CreateAddressResponse, error) = civicrm.CreateAddress

func getContact(donation *iraiser.Donation) (int, error) {
	query := civicrm.GetContactQuery{
		Mail: donation.Donator.Mail,
	}
	searchResult, err := contactGetter(&query);
	if err != nil {
		return 0, err
	}
	if searchResult.Count == 0 {
		return 0, &NoSuchContactError{Mail: donation.Donator.Mail}
	}
	if searchResult.Count > 1 {
		return 0, &TooManyContactsError{Mail: donation.Donator.Mail, Count: searchResult.Count}
	}
	return searchResult.Id, nil
}

func createContact(donation *iraiser.Donation) (int, error) {
	contact := civicrm.Contact{
		ContactType: internal.Config.ContactTypeName,
		Mail: donation.Donator.Mail,
		FirstName: donation.Donator.FirstName,
		LastName: donation.Donator.LastName,
		Pseudo: donation.Donator.Pseudo,
		Source: internal.Config.ContactSourceName,
	}
	contactResp, err := contactCreator(&contact)
	if err != nil {
		return 0, err
	}
	address := civicrm.Address{
		ContactId: contactResp.Id,
		LocationTypeId: internal.Config.LocationTypeId,
		StreetAddress: donation.Donator.StreetAddress,
		City: donation.Donator.City,
		PostalCode: donation.Donator.PostalCode,
		Country: donation.Donator.Country,
		StreetParsing: 1,
	}
	_, err = addressCreator(&address)
	return contactResp.Id, err
}
