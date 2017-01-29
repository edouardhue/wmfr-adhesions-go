package civicrm

type MembershipPayment struct {
	ContributionId int `json:"contribution_id" binding:"required"`
	MembershipId   int `json:"membership_id" binding:"required"`
}

type CreateMembershipPaymentResponse struct {
	StatusResponse
}

func (c *CiviCRM) CreateMembershipPayment(payment *MembershipPayment) (response *CreateMembershipPaymentResponse, _ error) {
	response = &CreateMembershipPaymentResponse{}
	req, err := c.buildQuery("MembershipPayment", "create", payment)
	if err != nil {
		return nil, err
	}
	return response, c.query(response, req)
}
