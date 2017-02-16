package civicrm

type Contact struct {
	ContactType        string `json:"contact_type"`
	Mail               string `json:"email"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Pseudo             string  `json:"pseudo"`
	Source             string `json:"source"`
}

type GetContactQuery struct {
	Mail string `json:"email"`
}

type GetContactResponse struct {
	StatusResponse
}

type CreateContactResponse struct {
	StatusResponse
}

func GetContact(query *GetContactQuery) (response *GetContactResponse, _ error) {
	response = &GetContactResponse{}
	req, err := buildQuery("Contact", "get", query)
	if err != nil {
		return nil, err
	}
	return response, execute(response, req)
}

func CreateContact(contact *Contact) (response *CreateContactResponse, _ error) {
	response = &CreateContactResponse{}
	req, err := buildQuery("Contact", "create", contact)
	if err != nil {
		return nil, err
	}
	return response, execute(response, req)
}
