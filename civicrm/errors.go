package civicrm

import "fmt"

type NoSuchMembershipError struct {
	typeId int
	r *GetMembershipResponse
}

func (e NoSuchMembershipError) Error() string {
	return fmt.Sprintf("no membership of type %d in %s", e.typeId, e.r.Values)
}

