package memberships

import (
	"github.com/wikimedia-france/wmfr-adhesions/iraiser"
	"github.com/wikimedia-france/wmfr-adhesions/civicrm"
	"github.com/wikimedia-france/wmfr-adhesions/internal"
)


func RecordMembership(donation *iraiser.Donation) error {
	contactId, err := getContact(donation);
	if err != nil {
		switch err.(type) {
		case NoSuchContactError:
			contactId, err := createContact(donation)
			if err != nil {
				return err
			}
			return recordNewMembership(donation, contactId)
		default:
			return err
		}
	}
	err = recordMembershipRenewal(donation, contactId);
	if err != nil {
		switch err.(type) {
		case NoSuitableMembershipError:
			return recordNewMembership(donation, contactId)
		default:
			return err
		}
	}
	return nil
}

func recordMembershipRenewal(donation *iraiser.Donation, contactId int) error {
	memberships, err := civicrm.GetMembership(&civicrm.GetMembershipQuery{ContactId: contactId})
	if err != nil {
		return err
	}
	membership, err := memberships.FindFirstByType(internal.Config.MembershipTypeId)
	if err != nil {
		return &NoSuitableMembershipError{Mail: donation.Donator.Mail, ExpectedMembershipTypeId: internal.Config.MembershipTypeId}
	}
	err = recordContribution(donation, membership)
	if err != nil {
		return err
	}
	return renewMembership(donation, membership)
}

func recordNewMembership(donation *iraiser.Donation, contactId int) error {
	membership, err := createMembership(donation, contactId)
	if err != nil {
		return err
	}
	return recordContribution(donation, membership)
}

func createMembership(donation *iraiser.Donation, contactId int) (membership *civicrm.Membership, err error) {
	membership = &civicrm.Membership{
		ContactId: contactId,
		EndDate: civicrm.Date{},
		StartDate: civicrm.Date{Time: donation.ValidationDate},
		JoinDate: civicrm.Date{Time: donation.ValidationDate},
		StatusOverride: 1,
		StatusId: internal.Config.MembershipStatusId,
		Terms: 1,
		CampaignId: internal.Config.CampaignId,
	}
	_, err = civicrm.CreateMembership(membership)
	return
}

func renewMembership(donation *iraiser.Donation, membership *civicrm.Membership) error {
	membership.EndDate = civicrm.Date{}
	membership.StartDate = civicrm.Date{Time: donation.ValidationDate}
	membership.StatusOverride = 1
	membership.StatusId = internal.Config.MembershipStatusId
	membership.Terms = 1
	membership.CampaignId = internal.Config.CampaignId
	_, err := civicrm.CreateMembership(membership)
	return err
}
