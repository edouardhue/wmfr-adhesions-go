package civicrm

import (
	"math/big"
	"time"
)

type Config struct {
	URL     string `yaml:"url"`
	SiteKey string `yaml:"siteKey"`
	UserKey string `yaml:"userKey"`
}

type Response interface {
	Success() bool
	GetErrorMessage() string
}

type ResponseError struct {
	Message string
}

func (e ResponseError) Error() string {
	return e.Message
}

type SearchContactQuery struct {
	EMail string `json:"email"`
}

type SearchContactResponse struct {
	IsError      int `json:"is_error" binding:"required"`
	ErrorMessage string `json:"error_message"`
	Version      int `json:"version" binding:"required"`
	Count        int `json:"count" binding:"required"`
	ContactId    int `json:"id"`
}

func (r *SearchContactResponse) Success() bool {
	return r.IsError == 0
}

func (r *SearchContactResponse) GetErrorMessage() string {
	return r.ErrorMessage
}

type ListMembershipsQuery struct {
	ContactId int `json:"contact_id"`
}

type ListMembershipsResponse struct {
	IsError      int `json:"is_error" binding:"required"`
	ErrorMessage string `json:"error_message"`
	Version      int `json:"version" binding:"required"`
	Count        int `json:"count" binding:"required"`
	Values       map[int]Membership `json:"values"`
}

type Membership struct {
	Id               int `json:"id,string" binding:"required"`
	ContactId        int `json:"contact_id,string" binding:"required"`
	MembershipTypeId int `json:"membership_type_id,string" binding:"required"`
	JoinDate         string `json:"join_date" binding:"required"`
	StartDate        string `json:"start_date" binding:"required"`
	EndDate          string `json:"end_date" binding:"required"`
	StatusId         int `json:"status_id,string" binding:"required"`
	IsOverride       int `json:"is_override,string"`
}

func (r *ListMembershipsResponse) Success() bool {
	return r.IsError == 0
}

func (r *ListMembershipsResponse) GetErrorMessage() string {
	return r.ErrorMessage
}

type Contribution struct {
	ContactId           int `json:"contact_id" binding:"required"`
	FinancialTypeId     int `json:"financial_type_id" binding:"required"`
	TotalAmount         *big.Rat `json:"total_amount" binding:"required"`
	Currency            string `json:"currency" binding:"required"`
	PaymentInstrumentId int `json:"payment_instrument_id" binding:"required"`
	ReceiveDate         time.Time `json:"receive_date" binding:"required"`
	TrxnId              string `json:"trxn_id" binding:"required"`
	InvoiceId           string `json:"invoice_id" binding:"required"`
	Source              string `json:"source" binding:"required"`
	CampaignId          int `json:"campaign_id" binding:"required"`
	StatusId            int `json:"status_id" binding:"required"`
}

type CreateContributionResponse struct {
	IsError      int `json:"is_error" binding:"required"`
	ErrorMessage string `json:"error_message"`
	Version      int `json:"version" binding:"required"`
}

func (r *CreateContributionResponse) Success() bool {
	return r.IsError == 0
}

func (r *CreateContributionResponse) GetErrorMessage() string {
	return r.ErrorMessage
}