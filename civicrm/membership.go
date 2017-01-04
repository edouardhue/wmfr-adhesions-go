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

func (c *CiviCRM) GetMembership(query *GetMembershipQuery) (response *GetMembershipResponse, _ error) {
	response = &GetMembershipResponse{}
	if req, err := c.buildQuery("Membership", "get", query); err != nil {
		return nil, err
	} else {
		return response, c.query(response, req)
	}
}

func (c *CiviCRM) CreateMembership(membership *Membership) (response *CreateMembershipResponse, _ error) {
	response = &CreateMembershipResponse{}
	if req, err := c.buildQuery("Membership", "create", membership); err != nil {
		return nil, err
	} else {
		return response, c.query(response, req)
	}
}
