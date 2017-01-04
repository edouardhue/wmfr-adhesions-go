package memberships

import "fmt"

type NoCommonMembershipError struct {
	Mail string
}

func (e *NoCommonMembershipError) Error() string {
	return fmt.Sprintf("%s - no common membership", e.Mail)
}
