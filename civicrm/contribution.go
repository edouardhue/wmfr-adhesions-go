package civicrm

import (
	"math/big"
	"time"
)

type Contribution struct {
	ContactId           int `json:"contact_id,string" binding:"required"`
	FinancialTypeId     int `json:"financial_type_id,string" binding:"required"`
	TotalAmount         *big.Rat `json:"total_amount" binding:"required"`
	Currency            string `json:"currency" binding:"required"`
	PaymentInstrumentId int `json:"payment_instrument_id,string" binding:"required"`
	ReceiveDate         time.Time `json:"receive_date" binding:"required"`
	TrxnId              string `json:"trxn_id" binding:"required"`
	InvoiceId           string `json:"invoice_id" binding:"required"`
	Source              string `json:"source" binding:"required"`
	CampaignId          int `json:"campaign_id,string" binding:"required"`
	StatusId            int `json:"status_id,string" binding:"required"`
}

type CreateContributionResponse struct {
	StatusResponse
}

func (c *CiviCRM) CreateContribution(contribution *Contribution) (response *CreateContributionResponse, _ error) {
	response = &CreateContributionResponse{}
	if req, err := c.buildQuery("Contribution", "create", contribution); err != nil {
		return nil, err
	} else {
		return response, c.query(response, req)
	}
}
