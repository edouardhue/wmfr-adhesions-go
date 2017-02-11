package iraiser

import (
	"testing"
	"encoding/json"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestDonationJsonBinding(t *testing.T) {
	bytes, err := ioutil.ReadFile("sample_donation.json")
	if err != nil {
		t.Fatal(err)
	}
	donation := []Donation{}
	if err := json.Unmarshal(bytes, &donation); err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, donation)
	assert.Equal(t, "john@doe.fr", donation[0].Donator.Mail)
	assert.Equal(t, "John-Jérémy", donation[0].Donator.FirstName)
	assert.Equal(t, "Doe", donation[0].Donator.LastName)
	assert.Equal(t, "Jdoe", donation[0].Donator.Pseudo)
	assert.Equal(t, "25 rue du port", donation[0].Donator.StreetAddress)
	assert.Equal(t, "Léon", donation[0].Donator.City)
	assert.Equal(t, "27400", donation[0].Donator.PostalCode)
	assert.Equal(t, "FR", donation[0].Donator.Country)
	assert.Equal(t, "", donation[0].Campaign.AffectationCode)
	assert.Equal(t, "", donation[0].Campaign.OriginCode)
	assert.Equal(t, "card", donation[0].Payment.Mode)
	assert.Equal(t, "9533016784", donation[0].Payment.GatewayId)
	assert.Equal(t, 1200, donation[0].Amount)
	assert.Equal(t, "EUR", donation[0].Currency)
	assert.Equal(t, "WIKI20170210-1298-1387", donation[0].Reference)
	assert.Equal(t, time.Date(2016, time.October, 28, 14, 18, 43, 0, time.Local), donation[0].ValidationDate)
}