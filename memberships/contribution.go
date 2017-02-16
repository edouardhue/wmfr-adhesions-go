package memberships

import (
	"github.com/wikimedia-france/wmfr-adhesions/iraiser"
	"github.com/wikimedia-france/wmfr-adhesions/civicrm"
	"math/big"
	"github.com/wikimedia-france/wmfr-adhesions/internal"
)

var contributionCreator func(*civicrm.Contribution) (*civicrm.CreateContributionResponse, error) = civicrm.CreateContribution
var membershipPaymentCreator func(*civicrm.MembershipPayment) (*civicrm.CreateMembershipPaymentResponse, error) = civicrm.CreateMembershipPayment

func recordContribution(donation *iraiser.Donation, membership *civicrm.Membership) (*civicrm.Contribution, error) {
	contribution := civicrm.Contribution{
		ContactId: membership.ContactId,
		FinancialTypeId: internal.Config.MembershipFinancialTypeId,
		TotalAmount: big.NewRat(int64(donation.Amount), 100),
		Currency: donation.Currency,
		PaymentInstrumentId: internal.Config.PaymentInstruments[donation.Payment.Mode],
		ReceiveDate: donation.ValidationDate,
		TrxnId: donation.Payment.GatewayId,
		InvoiceId: donation.Reference,
		Source: internal.Config.ContributionSourceName,
		CampaignId: internal.Config.CampaignId,
		StatusId: internal.Config.ContributionStatusId,
	}
	createReponse, err := contributionCreator(&contribution)
	if err != nil {
		return &civicrm.Contribution{}, err
	}
	payment := civicrm.MembershipPayment{
		ContributionId: createReponse.Id,
		MembershipId: membership.Id,
	}
	_, err = membershipPaymentCreator(&payment)
	return &contribution, err
}
