package memberships

import (
	"github.com/wikimedia-france/wmfr-adhesions/iraiser"
	"github.com/wikimedia-france/wmfr-adhesions/civicrm"
	"github.com/wikimedia-france/wmfr-adhesions/internal"
)

const StatusOverride = 1
const Terms = 1

var membershipGetter func(*civicrm.GetMembershipQuery) (*civicrm.GetMembershipResponse, error) = civicrm.GetMembership
var membershipCreator func(*civicrm.Membership) (*civicrm.CreateMembershipResponse, error) = civicrm.CreateMembership

func RecordMembership(donation *iraiser.Donation) (*civicrm.Membership, error) {
	contactId, err := getContact(donation);
	if err != nil {
		switch err.(type) {
		case NoSuchContactError:
			return recordNewContactMembership(donation)
		default:
			return &civicrm.Membership{}, err
		}
	}
	membership, err := recordMembershipRenewal(donation, contactId);
	if err != nil {
		switch err.(type) {
		case NoSuitableMembershipError:
			return recordNewMembership(donation, contactId)
		default:
			return &civicrm.Membership{}, err
		}
	}
	return membership, nil
}

func recordNewContactMembership(donation *iraiser.Donation) (*civicrm.Membership, error) {
	contactId, err := createContact(donation)
	if err != nil {
		return &civicrm.Membership{}, err
	}
	return recordNewMembership(donation, contactId)
}

func recordMembershipRenewal(donation *iraiser.Donation, contactId int) (*civicrm.Membership, error) {
	memberships, err := membershipGetter(&civicrm.GetMembershipQuery{ContactId: contactId})
	if err != nil {
		return &civicrm.Membership{}, err
	}
	membership, err := memberships.FindFirstByType(internal.Config.MembershipTypeId)
	if err != nil {
		return &civicrm.Membership{}, &NoSuitableMembershipError{Mail: donation.Donator.Mail, ExpectedMembershipTypeId: internal.Config.MembershipTypeId}
	}
	_, err = recordContribution(donation, membership)
	if err != nil {
		return &civicrm.Membership{}, err
	}
	return renewMembership(donation, membership)
}

func recordNewMembership(donation *iraiser.Donation, contactId int) (*civicrm.Membership, error) {
	membership, err := createMembership(donation, contactId)
	if err != nil {
		return &civicrm.Membership{}, err
	}
	_, err = recordContribution(donation, membership)
	return &civicrm.Membership{}, err
}

func createMembership(donation *iraiser.Donation, contactId int) (*civicrm.Membership, error) {
	membership := &civicrm.Membership{
		ContactId: contactId,
		EndDate: civicrm.Date{},
		StartDate: civicrm.Date{Time: donation.ValidationDate},
		JoinDate: civicrm.Date{Time: donation.ValidationDate},
		StatusOverride: StatusOverride,
		StatusId: internal.Config.MembershipStatusId,
		Terms: Terms,
		CampaignId: internal.Config.CampaignId,
	}
	resp, err := membershipCreator(membership)
	membership.Id = resp.Id
	return membership, err
}

func renewMembership(donation *iraiser.Donation, membership *civicrm.Membership) (*civicrm.Membership, error) {
	membership.EndDate = civicrm.Date{}
	membership.StartDate = civicrm.Date{Time: donation.ValidationDate}
	membership.StatusOverride = StatusOverride
	membership.StatusId = internal.Config.MembershipStatusId
	membership.Terms = Terms
	membership.CampaignId = internal.Config.CampaignId
	_, err := membershipCreator(membership)
	return membership, err
}
