package memberships

import (
	"testing"
	"github.com/wikimedia-france/wmfr-adhesions/iraiser"
	"github.com/wikimedia-france/wmfr-adhesions/civicrm"
	"github.com/stretchr/testify/assert"
	"github.com/wikimedia-france/wmfr-adhesions/internal"
	"math/rand"
	"math/big"
	"time"
)

func TestRecordContribution(t *testing.T) {
	financialTypeId := rand.Int()
	cardPaymentInstrumentId := rand.Int()
	contributionSourceName := "test"
	campaignId := rand.Int()
	statusId := rand.Int()
	validationDate := time.Now()
	amount := rand.Int()

	internal.Config.MembershipFinancialTypeId = financialTypeId
	internal.Config.PaymentInstruments = map[string]int {
		"card": cardPaymentInstrumentId,
	}
	internal.Config.ContributionSourceName = contributionSourceName
	internal.Config.CampaignId = campaignId
	internal.Config.ContributionStatusId = statusId

	donation := iraiser.Donation{
		Donator: iraiser.Donator{
			Mail: "test@example.org",
			FirstName: "First",
			LastName: "Last",
			Pseudo: "Nick",
			StreetAddress: "Address",
			City: "City",
			PostalCode: "12345",
			Country: "CY",
		},
		Campaign: iraiser.Campaign{
			AffectationCode: "Aff",
			OriginCode: "Ori",
		},
		Payment: iraiser.Payment{
			Mode: "card",
			GatewayId: "6789",
		},
		Amount: amount,
		Currency: "CUR",
		Reference: "abcd",
		ValidationDate: validationDate,
	}
	membership := civicrm.Membership{}

	contributionCreator = func(*civicrm.Contribution) (*civicrm.CreateContributionResponse, error) {
		return &civicrm.CreateContributionResponse{}, nil
	}
	membershipPaymentCreator = func(*civicrm.MembershipPayment) (*civicrm.CreateMembershipPaymentResponse, error) {
		return &civicrm.CreateMembershipPaymentResponse{}, nil
	}

	contribution, err := recordContribution(&donation, &membership)

	assert.NoError(t, err, "Record result")
	assert.Equal(t, membership.ContactId, contribution.ContactId, "Contact id is propagated")
	assert.Equal(t, financialTypeId, contribution.FinancialTypeId, "Financial type id is propagated")
	assert.Equal(t, big.NewRat(int64(amount), 100), contribution.TotalAmount, "Amount is propagated")
	assert.Equal(t, donation.Currency, contribution.Currency, "Currency is propagated")
	assert.Equal(t, cardPaymentInstrumentId, contribution.PaymentInstrumentId, "Payment instrument id is retrieved")
	assert.Equal(t, donation.ValidationDate, contribution.ReceiveDate, "Validation date is propagatated")
	assert.Equal(t, donation.Payment.GatewayId, contribution.TrxnId, "Gateway id is propagated")
	assert.Equal(t, donation.Reference, contribution.InvoiceId, "Reference is propagated")
	assert.Equal(t, contributionSourceName, contribution.Source, "Contribution source is propagated")
	assert.Equal(t, campaignId, contribution.CampaignId, "Campaign id is propagated")
	assert.Equal(t, statusId, contribution.StatusId, "Status id is propagated")
}

func TestRecordContributionWithCrmContributionError(t *testing.T) {
	donation := iraiser.Donation{}
	membership := civicrm.Membership{}

	contributionCreator = func(*civicrm.Contribution) (*civicrm.CreateContributionResponse, error) {
		return &civicrm.CreateContributionResponse{}, assert.AnError
	}
	membershipPaymentCreator = func(*civicrm.MembershipPayment) (*civicrm.CreateMembershipPaymentResponse, error) {
		t.Fatal("Should not attempt to create a membership payment")
		return &civicrm.CreateMembershipPaymentResponse{}, nil
	}

	_, err := recordContribution(&donation, &membership)

	assert.Error(t, err)
}

func TestRecordContributionWithCrmMembershipPaymentError(t *testing.T) {
	donation := iraiser.Donation{}
	membership := civicrm.Membership{}

	contributionCreator = func(*civicrm.Contribution) (*civicrm.CreateContributionResponse, error) {
		return &civicrm.CreateContributionResponse{}, nil
	}
	membershipPaymentCreator = func(*civicrm.MembershipPayment) (*civicrm.CreateMembershipPaymentResponse, error) {
		return &civicrm.CreateMembershipPaymentResponse{}, assert.AnError
	}

	_, err := recordContribution(&donation, &membership)

	assert.Error(t, err)
}
