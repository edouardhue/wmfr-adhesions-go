package internal

import (
	"os"
	"gopkg.in/yaml.v2"
	"fmt"
)

var Config *Configuration

type Configuration struct {
	CiviCRM                   CiviCRMConfiguration `yaml:"civicrm"`
	IRaiser                   IRaiserConfiguration `yaml:"iRaiser"`
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

type CiviCRMConfiguration struct {
	URL     string `yaml:"url"`
	SiteKey string `yaml:"siteKey"`
	UserKey string `yaml:"userKey"`
}

type IRaiserConfiguration struct {
	SecureKey string `yaml:"secureKey"`
}


func init() {
	var location, exists = os.LookupEnv("CONFIG_LOCATION")
	if !exists {
		location = "./adhesions.yaml"
	}
	fileinfo, err := os.Stat(location)
	if err != nil {
		fmt.Println("Cannot use configuration file, using defaults.", err)
		Config = &Configuration{}
		Config.CiviCRM = CiviCRMConfiguration{}
		Config.IRaiser = IRaiserConfiguration{}
		return
	}
	filesize := fileinfo.Size()
	fp, err := os.Open(location)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	buf := make([]byte, filesize)
	fp.Read(buf)
	if err = yaml.Unmarshal(buf, &Config); err != nil {
		panic(err)
	}
}