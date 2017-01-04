package iraiser

import (
	"time"
	"crypto/md5"
	"bytes"
)

type Config struct {
	SecureKey string `yaml:"secureKey"`
}

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
	Pseudo        string `json:"reserved_pseudo" binding:"required"`
	StreetAddress string `json:"address1" binding:"required"`
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

type SecureHeader struct {
	Login     string
	Timestamp string
	Token     []byte
}

type IRaiser struct {
	config *Config
}

func NewIRaiser(config *Config) *IRaiser {
	return &IRaiser{
		config: config,
	}
}

func (i *IRaiser) Verify(h *SecureHeader) bool {
	var expected = md5.Sum([]byte(h.Login + i.config.SecureKey + h.Timestamp))
	return bytes.Equal(expected[:], h.Token)
}