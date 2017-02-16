package memberships

import (
	"github.com/wikimedia-france/wmfr-adhesions/iraiser"
	"github.com/wikimedia-france/wmfr-adhesions/civicrm"
	"math/big"
	"github.com/wikimedia-france/wmfr-adhesions/internal"
)

func recordContribution(donation *iraiser.Donation, membership *civicrm.Membership) error {
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
	createReponse, err := civicrm.CreateContribution(&contribution)
	if err != nil {
		return err
	}
	payment := civicrm.MembershipPayment{
		ContributionId: createReponse.Id,
		MembershipId: membership.Id,
	}
	_, err = civicrm.CreateMembershipPayment(&payment)
	return err
}
