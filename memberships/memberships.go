package memberships

import (
	"github.com/edouardhue/wmfr-adhesions/iraiser"
	"github.com/edouardhue/wmfr-adhesions/civicrm"
	"math/big"
	"net/http"
)

type Memberships struct {
	config *Config
	crm    *civicrm.CiviCRM
}

func NewMemberships(config *Config) *Memberships {
	return &Memberships{
		config: config,
		crm: civicrm.NewCiviCRM(&config.CiviCRM, http.DefaultClient),
	}
}

func (m *Memberships) RecordMembership(donation *iraiser.Donation) error {
	if searchResult, err := m.crm.GetContact(&civicrm.ContactQuery{EMail: donation.Donator.Mail}); err != nil {
		return err
	} else if searchResult.Count == 1 {
		return m.updateMembership(donation, searchResult.Id)
	} else {
		return &NoSuchContactError{donation.Donator.Mail}
	}
}

func (m *Memberships) updateMembership(donation *iraiser.Donation, contactId int) error {
	if memberships, err := m.crm.GetMembership(&civicrm.MembershipQuery{ContactId: contactId}); err != nil {
		return err
	} else if commonMembership := m.findCommonMembership(memberships); commonMembership != nil {
		if err := m.recordContribution(donation, commonMembership); err != nil {
			return err
		} else {
			return m.renewMembership(donation, commonMembership)
		}
	} else {
		return &NoCommonMembershipError{donation.Donator.Mail}
	}
}

func (m *Memberships) findCommonMembership(memberships *civicrm.GetMembershipResponse) *civicrm.Membership {
	for _, membership := range memberships.Values {
		if membership.MembershipTypeId == m.config.MembershipTypeId {
			return &membership
		}
	}
	return nil
}

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

func (m *Memberships) renewMembership(donation *iraiser.Donation, membership *civicrm.Membership) error {
	membership.EndDate = civicrm.Date{}
	membership.StartDate = civicrm.Date{Time: donation.ValidationDate}
	membership.StatusOverride = 1
	membership.StatusId = m.config.MembershipStatusId
	membership.Terms = 1
	membership.CampaignId = m.config.CampaignId
	_, err := m.crm.CreateMembership(membership)
	return err
}