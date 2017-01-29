package memberships

import (
	"github.com/wikimedia-france/wmfr-adhesions/iraiser"
	"github.com/wikimedia-france/wmfr-adhesions/civicrm"
	"net/http"
)

type Memberships struct {
	config *Config
	crm    *civicrm.CiviCRM
}

func NewMemberships(config *Config) *Memberships {
	return &Memberships{
		config: config,
		crm: civicrm.NewCiviCRM(&config.CiviCRM, http.DefaultClient),
	}
}

func (m *Memberships) RecordMembership(donation *iraiser.Donation) error {
	contactId, err := m.getContact(donation);
	if err != nil {
		switch err.(type) {
		case NoSuchContactError:
			contactId, err := m.createContact(donation)
			if err != nil {
				return err
			}
			return m.recordNewMembership(donation, contactId)
		default:
			return err
		}
	}
	err = m.recordMembershipRenewal(donation, contactId);
	if err != nil {
		switch err.(type) {
		case NoSuitableMembershipError:
			return m.recordNewMembership(donation, contactId)
		default:
			return err
		}
	}
	return nil
}