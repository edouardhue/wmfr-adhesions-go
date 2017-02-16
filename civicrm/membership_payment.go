package civicrm

type MembershipPayment struct {
	ContributionId int `json:"contribution_id" binding:"required"`
	MembershipId   int `json:"membership_id" binding:"required"`
}

type CreateMembershipPaymentResponse struct {
	StatusResponse
}

func CreateMembershipPayment(payment *MembershipPayment) (response *CreateMembershipPaymentResponse, _ error) {
	response = &CreateMembershipPaymentResponse{}
	req, err := buildQuery("MembershipPayment", "create", payment)
	if err != nil {
		return nil, err
	}
	return response, execute(response, req)
}
