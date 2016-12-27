package civicrm

import "math/big"

type ContactSearchResponse struct {
	IsError   int `json:"is_error" binding:"required"`
	Version   int `json:"version" binding:"required"`
	Count     int `json:"count" binding:"required"`
	ContactId int `json:"id"`
}

type ListMembershipsResponse struct {
	IsError int `json:"is_error" binding:"required"`
	Version int `json:"version" binding:"required"`
	Count int `json:"count" binding:"required"`
	Values map[int]Membership `json:"values"`
}

type Membership struct {
	Id int `json:"id,string" binding:"required"`
	ContactId int `json:"contact_id,string" binding:"required"`
	MembershipTypeId int `json:"membership_type_id,string" binding:"required"`
	JoinDate string `json:"join_date" binding:"required"`
	StartDate string `json:"start_date" binding:"required"`
	EndDate string `json:"end_date" binding:"required"`
	StatusId int `json:"status_id,string" binding:"required"`
	IsOverride int `json:"is_override,string"`
}

type Contribution struct {
	FinancialTypeId int
	TotalAmount big.Float
	ContactId int
}

type CreateContributionResponse struct {
	IsError   int `json:"is_error" binding:"required"`
	Version   int `json:"version" binding:"required"`
}