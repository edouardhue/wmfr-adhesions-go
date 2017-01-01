package memberships

import "fmt"

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
