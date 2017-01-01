package memberships

import (
	"github.com/edouardhue/wmfr-adhesions/civicrm"
	"github.com/edouardhue/wmfr-adhesions/iraiser"
)

type Config struct {
	CiviCRM                   civicrm.Config `yaml:"civicrm"`
	IRaiser                   iraiser.Config `yaml:"iRaiser"`
	PaymentInstruments        map[string]int `yaml:"paymentInstruments"`
	CampaignId                int `yaml:"campaignId"`
	MembershipTypeId          int `yaml:"membershipTypeId"`
	MembershipFinancialTypeId int `yaml:"membershipFinancialTypeId"`
	MembershipStatusId        int `yaml:"membershipStatusId"`
	ContributionStatusId      int `yaml:"contributionStatusId"`
	ContributionSourceName    string `yaml:"contributionSourceName"`
}