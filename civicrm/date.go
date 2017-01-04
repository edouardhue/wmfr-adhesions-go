package civicrm

import (
	"time"
	"errors"
)

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
