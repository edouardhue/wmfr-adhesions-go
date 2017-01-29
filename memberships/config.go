package memberships

import (
	"github.com/wikimedia-france/wmfr-adhesions/civicrm"
	"github.com/wikimedia-france/wmfr-adhesions/iraiser"
)

type Config struct {
	CiviCRM                   civicrm.Config `yaml:"civicrm"`
	IRaiser                   iraiser.Config `yaml:"iRaiser"`
	PaymentInstruments        map[string]int `yaml:"paymentInstruments"`
	CampaignId                int `yaml:"campaignId"`
	ContactTypeName           string `yaml:"contactTypeName"`
	ContactSourceName         string `yaml:"contactSourceName"`
	LocationTypeId            int `yaml:"locationTypeId"`
	MembershipTypeId          int `yaml:"membershipTypeId"`
	MembershipFinancialTypeId int `yaml:"membershipFinancialTypeId"`
	MembershipStatusId        int `yaml:"membershipStatusId"`
	ContributionStatusId      int `yaml:"contributionStatusId"`
	ContributionSourceName    string `yaml:"contributionSourceName"`
	Log                       string `yaml:"log"`
}