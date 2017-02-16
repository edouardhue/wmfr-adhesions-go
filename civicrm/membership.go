package civicrm

type Membership struct {
	Id               int `json:"id,string"`
	ContactId        int `json:"contact_id,string"`
	CampaignId       int `json:"campaign_id,string"`
	MembershipTypeId int `json:"membership_type_id,string"`
	JoinDate         Date `json:"join_date"`
	StartDate        Date `json:"start_date"`
	EndDate          Date `json:"end_date,omitempty"`
	StatusId         int `json:"status_id,string"`
	StatusOverride   int `json:"is_override,string"`
	Terms            int `json:"num_terms,string"`
}

type GetMembershipQuery struct {
	ContactId int `json:"contact_id"`
}

type GetMembershipResponse struct {
	StatusResponse
	Values       map[int]Membership `json:"values"`
}

type CreateMembershipResponse struct {
	StatusResponse
}

func GetMembership(query *GetMembershipQuery) (response *GetMembershipResponse, _ error) {
	response = &GetMembershipResponse{}
	req, err := buildQuery("Membership", "get", query)
	if err != nil {
		return nil, err
	}
	return response, execute(response, req)
}

func CreateMembership(membership *Membership) (response *CreateMembershipResponse, _ error) {
	response = &CreateMembershipResponse{}
	req, err := buildQuery("Membership", "create", membership)
	if err != nil {
		return nil, err
	}
	return response, execute(response, req)
}

func (r *GetMembershipResponse) FindFirstByType(typeId int) (*Membership, error) {
	for _, membership := range r.Values {
		if membership.MembershipTypeId == typeId {
			return &membership, nil
		}
	}
	return &Membership{}, NoSuchMembershipError{typeId, r}
}