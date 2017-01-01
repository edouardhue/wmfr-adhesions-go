package memberships

import (
	"github.com/edouardhue/wmfr-adhesions/civicrm"
	"github.com/edouardhue/wmfr-adhesions/iraiser"
)

type Config struct {
	CiviCRM                   civicrm.Config `yaml:"civicrm"`
	IRaiser                   iraiser.Config `yaml:"iRaiser"`
	PaymentInstruments        map[string]int `yaml:"paymentInstruments"`
	MembershipTypeId          int `yaml:"membershipTypeId"`
	MembershipFinancialTypeId int `yaml:"membershipFinancialTypeId"`
	MembershipCampaignId      int `yaml:"membershipCampaignId"`
	ContributionStatusId      int `yaml:"contributionStatusId"`
	ContributionSourceName    string `yaml:"contributionSourceName"`
}