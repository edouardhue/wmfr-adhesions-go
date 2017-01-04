package civicrm

type Contact struct {
	ContactType        int `json:"contact_type,string"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Pseudo             string  `json:"pseudo"`
	Source             string `json:"source"`
	EMail              string `json:"email"`
	ExternalIdentifier string `json:"external_identifier"`
}

type GetContactQuery struct {
	EMail string `json:"email"`
}

type GetContactResponse struct {
	StatusResponse
}

type CreateContactResponse struct {
	StatusResponse
}

func (c *CiviCRM) GetContact(query *GetContactQuery) (response *GetContactResponse, _ error) {
	response = &GetContactResponse{}
	if req, err := c.buildQuery("Contact", "get", query); err != nil {
		return nil, err
	} else {
		return response, c.query(response, req)
	}
}

func (c *CiviCRM) CreateContact(contact string) {

}
