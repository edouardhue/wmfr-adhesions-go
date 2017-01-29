package iraiser

import (
	"time"
)

type Donation struct {
	Donator        Donator `json:"donator" binding:"required"`
	Campaign       Campaign `json:"campaign" binding:"required"`
	Payment        Payment `json:"payment" binding:"required"`
	Amount         int `json:"amount" binding:"required"`
	Currency       string `json:"currency" binding:"required"`
	Reference      string `json:"reference" binding:"required"`
	ValidationDate time.Time `json:"validationDate" binding:"required"`
}

type Donator struct {
	Mail          string `json:"email" binding:"required"`
	FirstName     string `json:"firstName" binding:"required"`
	LastName      string `json:"lastName" binding:"required"`
	Pseudo        string `json:"pseudo" binding:"required"`
	StreetAddress string `json:"address" binding:"required"`
	City          string `json:"city" binding:"required"`
	PostalCode    string `json:"postcode" binding:"required"`
	Country       string `json:"country" binding:"required"`
}

type Campaign struct {
	AffectationCode string `json:"affectationCode" binding:"required"`
	OriginCode      string `json:"originCode" binding:"required"`
}

type Payment struct {
	Mode      string `json:"mode" binding:"required"`
	GatewayId string `json:"gatewayId" binding:"required"`
}