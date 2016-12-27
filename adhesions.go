package main

import (
	"fmt"
	"github.com/edouardhue/wmfr-adhesions/iraiser"
	"github.com/edouardhue/wmfr-adhesions/civicrm"
)

const COMMON_MEMBERSHIP_ID = 9

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

func updateOrCreateMembership(member iraiser.Member) error {
	if searchResult, err := civicrm.SearchContact(member.Mail) ; err != nil {
		return err
	} else if searchResult.Count == 1 {
		if memberships, err := civicrm.ListContactMemberships(searchResult.Id); err != nil {
			return err
		} else {
			if commonMembership := findCommonMembership(memberships) ; commonMembership != nil {
				return nil
			} else {
				return &NoCommonMembershipError{member.Mail}
			}
		}
	} else {
		return &NoSuchContactError{member.Mail}
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