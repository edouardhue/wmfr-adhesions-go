package civicrm

import (
	"time"
	"errors"
)

type MembershipQuery struct {
	ContactId int `json:"contact_id"`
}

type GetMembershipResponse struct {
	IsError      int `json:"is_error" binding:"required"`
	ErrorMessage string `json:"error_message"`
	Version      int `json:"version"`
	Count        int `json:"count"`
	Values       map[int]Membership `json:"values"`
}

type Membership struct {
	Id               int `json:"id,string" binding:"required"`
	ContactId        int `json:"contact_id,string"`
	CampaignId       int `json:"campaign_id,string"`
	MembershipTypeId int `json:"membership_type_id,string"`
	JoinDate         Date `json:"join_date"`
	StartDate        Date `json:"start_date"`
	EndDate          Date `json:"end_date,omitempty"`
	StatusId         int `json:"status_id,string"`
	StatusOverride   int `json:"is_override,string"`
	Terms            int `json:"num_terms,string"`
}

func (r *GetMembershipResponse) Success() bool {
	return r.IsError == 0
}

func (r *GetMembershipResponse) GetErrorMessage() string {
	return r.ErrorMessage
}

type CreateMembershipResponse struct {
	IsError      int `json:"is_error" binding:"required"`
	ErrorMessage string `json:"error_message"`
	Version      int `json:"version"`
	Id           int `json:"id"`
}

func (r *CreateMembershipResponse) Success() bool {
	return r.IsError == 0
}

func (r *CreateMembershipResponse) GetErrorMessage() string {
	return r.ErrorMessage
}


const DATE_FORMAT = "2006-01-02"

type Date struct {
	time.Time
}

// MarshalJSON implements the json.Marshaler interface.
func (t Date) MarshalJSON() ([]byte, error) {
	if y := t.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Date.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DATE_FORMAT)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, DATE_FORMAT)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *Date) UnmarshalJSON(data []byte) error {
	// Fractional seconds are handled implicitly by Parse.
	var err error
	t.Time, err = time.Parse(`"`+DATE_FORMAT+`"`, string(data))
	return err
}