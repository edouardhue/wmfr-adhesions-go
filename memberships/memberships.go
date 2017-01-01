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

func (m *Memberships) RecordMembership(donation iraiser.Donation) error {
	if searchResult, err := m.crm.SearchContact(civicrm.SearchContactQuery{EMail: donation.Donator.Mail}); err != nil {
		return err
	} else if searchResult.Count == 1 {
		return m.updateMembership(donation, searchResult.ContactId)
	} else {
		return &NoSuchContactError{donation.Donator.Mail}
	}
}

func (m *Memberships) updateMembership(donation iraiser.Donation, contactId int) error {
	if memberships, err := m.crm.ListContactMemberships(civicrm.ListMembershipsQuery{ContactId: contactId}); err != nil {
		return err
	} else if commonMembership := m.findCommonMembership(memberships); commonMembership != nil {
		return m.recordContribution(donation, contactId)
	} else {
		return &NoCommonMembershipError{donation.Donator.Mail}
	}
}

func (m *Memberships) findCommonMembership(memberships *civicrm.ListMembershipsResponse) *civicrm.Membership {
	for _, membership := range memberships.Values {
		if membership.MembershipTypeId == m.config.MembershipTypeId {
			return &membership
		}
	}
	return nil
}

func (m *Memberships) recordContribution(donation iraiser.Donation, contactId int) error {
	contribution := civicrm.Contribution{
		ContactId: contactId,
		FinancialTypeId: m.config.MembershipFinancialTypeId,
		TotalAmount: big.NewRat(int64(donation.Amount), 100),
		Currency: donation.Currency,
		PaymentInstrumentId: m.config.PaymentInstruments[donation.Payment.Mode],
		ReceiveDate: donation.ValidationDate,
		TrxnId: donation.Payment.GatewayId,
		InvoiceId: donation.Reference,
		Source: m.config.ContributionSourceName,
		CampaignId: m.config.MembershipCampaignId,
		StatusId: m.config.ContributionStatusId,
	}
	_, err := m.crm.CreateContribution(contribution)
	return err
}