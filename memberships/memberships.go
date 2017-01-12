package memberships

import (
	"github.com/edouardhue/wmfr-adhesions/iraiser"
	"github.com/edouardhue/wmfr-adhesions/civicrm"
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
	if contactId, err := m.getContact(donation); err != nil {
		switch err.(type) {
		case NoSuchContactError:
			if contactId, err := m.createContact(donation); err != nil {
				return err
			} else {
				return m.recordNewMembership(donation, contactId)
			}
		default:
			return err
		}
	} else if err := m.recordMembershipRenewal(donation, contactId); err != nil {
		switch err.(type) {
		case NoSuitableMembershipError:
			return m.recordNewMembership(donation, contactId)
		default:
			return err
		}
	} else {
		return nil
	}
}