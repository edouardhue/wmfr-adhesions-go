package memberships

import (
	"github.com/wikimedia-france/wmfr-adhesions/iraiser"
	"github.com/wikimedia-france/wmfr-adhesions/civicrm"
)

func (m *Memberships) recordMembershipRenewal(donation *iraiser.Donation, contactId int) error {
	memberships, err := m.crm.GetMembership(&civicrm.GetMembershipQuery{ContactId: contactId})
	if err != nil {
		return err
	}
	membership, err := memberships.FindFirstByType(m.config.MembershipTypeId)
	if err != nil {
		return &NoSuitableMembershipError{Mail: donation.Donator.Mail, ExpectedMembershipTypeId: m.config.MembershipTypeId}
	}
	err = m.recordContribution(donation, membership)
	if err != nil {
		return err
	}
	return m.renewMembership(donation, membership)
}

func (m *Memberships) recordNewMembership(donation *iraiser.Donation, contactId int) error {
	membership, err := m.createMembership(donation, contactId)
	if err != nil {
		return err
	}
	return m.recordContribution(donation, membership)
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
