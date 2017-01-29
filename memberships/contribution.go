package memberships

import (
	"github.com/wikimedia-france/wmfr-adhesions/iraiser"
	"github.com/wikimedia-france/wmfr-adhesions/civicrm"
	"math/big"
)

func (m *Memberships) recordContribution(donation *iraiser.Donation, membership *civicrm.Membership) error {
	contribution := civicrm.Contribution{
		ContactId: membership.ContactId,
		FinancialTypeId: m.config.MembershipFinancialTypeId,
		TotalAmount: big.NewRat(int64(donation.Amount), 100),
		Currency: donation.Currency,
		PaymentInstrumentId: m.config.PaymentInstruments[donation.Payment.Mode],
		ReceiveDate: donation.ValidationDate,
		TrxnId: donation.Payment.GatewayId,
		InvoiceId: donation.Reference,
		Source: m.config.ContributionSourceName,
		CampaignId: m.config.CampaignId,
		StatusId: m.config.ContributionStatusId,
	}
	if createReponse, err := m.crm.CreateContribution(&contribution); err != nil {
		return err
	} else {
		payment := civicrm.MembershipPayment{
			ContributionId: createReponse.Id,
			MembershipId: membership.Id,
		}
		_, err := m.crm.CreateMembershipPayment(&payment)
		return err
	}
}
