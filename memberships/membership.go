package memberships

import (
	"github.com/edouardhue/wmfr-adhesions/iraiser"
	"github.com/edouardhue/wmfr-adhesions/civicrm"
)

func (m *Memberships) recordMembershipRenewal(donation *iraiser.Donation, contactId int) error {
	if memberships, err := m.crm.GetMembership(&civicrm.GetMembershipQuery{ContactId: contactId}); err != nil {
		return err
	} else if membership := memberships.FindFirstByType(m.config.MembershipTypeId) ; membership != nil {
		if err := m.recordContribution(donation, membership); err != nil {
			return err
		} else {
			return m.renewMembership(donation, membership)
		}
	} else {
		return &NoSuitableMembershipError{Mail: donation.Donator.Mail, ExpectedMembershipTypeId: m.config.MembershipTypeId}
	}
}

func (m *Memberships) recordNewMembership(donation *iraiser.Donation, contactId int) error {
	if membership, err := m.createMembership(donation, contactId); err != nil {
		return err
	} else {
		return m.recordContribution(donation, membership)
	}
}

func (m *Memberships) createMembership(donation *iraiser.Donation, contactId int) (membership *civicrm.Membership, err error) {
	membership = &civicrm.Membership{
		ContactId: contactId,
		EndDate: civicrm.Date{},
		StartDate: civicrm.Date{Time: donation.ValidationDate},
		JoinDate: civicrm.Date{Time: donation.ValidationDate},
		StatusOverride: 1,
		StatusId: m.config.MembershipStatusId,
		Terms: 1,
		CampaignId: m.config.CampaignId,
	}
	_, err = m.crm.CreateMembership(membership)
	return
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
