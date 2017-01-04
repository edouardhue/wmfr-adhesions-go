package civicrm

type Address struct {
	Id               int `json:"id,string,omitempty"`
	ContactId        int `json:"contact_id,string"`
	LocationTypeId   int `json:"location_type_id,string"`
	StreetAddress    string `json:"street_address"`
	StreetParsing    int `json:"street_parsing"`
	City             string `json:"city"`
	PostalCode       string `json:"postal_code"`
	Country          string `json:"country_id"`
}

type CreateAddressResponse struct {
	StatusResponse
}

func (c *CiviCRM) CreateAddress(address *Address) (response *CreateAddressResponse, _ error) {
	response = &CreateAddressResponse{}
	if req, err := c.buildQuery("Address", "create", address); err != nil {
		return nil, err
	} else {
		return response, c.query(response, req)
	}
}
