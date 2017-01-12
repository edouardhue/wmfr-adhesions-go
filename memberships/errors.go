package memberships

import "fmt"

type NoSuchContactError struct {
	Mail string
}

func (e NoSuchContactError) Error() string {
	return fmt.Sprintf("%s - no such contact", e.Mail)
}

type TooManyContactsError struct {
	Mail string
	Count int
}

func (e TooManyContactsError) Error() string {
	return fmt.Sprintf("%s - %d results", e.Mail, e.Count)
}

type NoSuitableMembershipError struct {
	Mail string
	ExpectedMembershipTypeId int
}

func (e NoSuitableMembershipError) Error() string {
	return fmt.Sprintf("%s - no membership of type %d", e.Mail, e.ExpectedMembershipTypeId)
}
