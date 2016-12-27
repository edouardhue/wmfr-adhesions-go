package main

import (
	"fmt"
	"github.com/edouardhue/wmfr-adhesions/iraiser"
	"github.com/edouardhue/wmfr-adhesions/civicrm"
	"math/big"
)

const COMMON_MEMBERSHIP_ID = 9

const MEMBERSHIP_FINANCIAL_TYPE_ID = 1

type NoSuchContactError struct {
	Mail string
}

func (e *NoSuchContactError) Error() string {
	return fmt.Sprintf("%s - no such contact", e.Mail)
}

type NoCommonMembershipError struct {
	Mail string
}

func (e *NoCommonMembershipError) Error() string {
	return fmt.Sprintf("%s - no common membership", e.Mail)
}

func recordMembership(member iraiser.Member) error {
	if searchResult, err := civicrm.SearchContact(member.Mail) ; err != nil {
		return err
	} else if searchResult.Count == 1 {
		return updateMembership(member, searchResult.ContactId)
	} else {
		return &NoSuchContactError{member.Mail}
	}
}

func updateMembership(member iraiser.Member, contactId int) error {
	if memberships, err := civicrm.ListContactMemberships(contactId); err != nil {
		return err
	} else if commonMembership := findCommonMembership(memberships) ; commonMembership != nil {
		return recordContribution(member.Amount, contactId)
	} else {
		return &NoCommonMembershipError{member.Mail}
	}
}

func findCommonMembership(memberships civicrm.ListMembershipsResponse) *civicrm.Membership {
	for _, m := range memberships.Values {
		if m.MembershipTypeId == COMMON_MEMBERSHIP_ID {
			return &m
		}
	}
	return nil
}

func recordContribution(amount big.Float, contactId int) error {
	contribution := civicrm.Contribution{
		FinancialTypeId: MEMBERSHIP_FINANCIAL_TYPE_ID,
		TotalAmount: amount,
		ContactId: contactId,
	}
	_, err := civicrm.CreateContribution(contribution)
	return err
}