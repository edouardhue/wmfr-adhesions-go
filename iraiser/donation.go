package iraiser

import (
	"time"
)

type Donation struct {
	Donator
	Campaign
	Payment
	Amount         int `json:"amount,string" binding:"required"`
	Currency       string `json:"currency" binding:"required"`
	Reference      string `json:"reference" binding:"required"`
	ValidationDate time.Time `json:"validationDate" binding:"required"`
}

type Donator struct {
	Mail          string `json:"donator.email" binding:"required"`
	FirstName     string `json:"donator.firstName" binding:"required"`
	LastName      string `json:"donator.lastName" binding:"required"`
	Pseudo        string `json:"donator.pseudo" binding:"required"`
	StreetAddress string `json:"donator.address" binding:"required"`
	City          string `json:"donator.city" binding:"required"`
	PostalCode    string `json:"donator.postcode" binding:"required"`
	Country       string `json:"donator.country" binding:"required"`
}

type Campaign struct {
	AffectationCode string `json:"campaign.affectationCode" binding:"required"`
	OriginCode      string `json:"campaign.originCode" binding:"required"`
}

type Payment struct {
	Mode      string `json:"payment.mode" binding:"required"`
	GatewayId string `json:"payment.gatewayId" binding:"required"`
}