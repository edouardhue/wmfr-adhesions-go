package civicrm

type MembershipPayment struct {
	ContributionId int `json:"contribution_id" binding:"required"`
	MembershipId   int `json:"membership_id" binding:"required"`
}

type CreateMembershipPaymentResponse struct {
	IsError      int `json:"is_error" binding:"required"`
	ErrorMessage string `json:"error_message"`
	Version      int `json:"version"`
	Count        int `json:"count"`
	Id           int `json:"id"`
}

func (r *CreateMembershipPaymentResponse) Success() bool {
	return r.IsError == 0
}

func (r *CreateMembershipPaymentResponse) GetErrorMessage() string {
	return r.ErrorMessage
}
