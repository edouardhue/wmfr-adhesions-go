package civicrm

import (
	"net/http"
	"encoding/json"
	"log"
	"net/url"
	"bytes"
)

type CiviCRM struct {
	client *http.Client
	config *Config
}

func NewCiviCRM(config *Config, client *http.Client) *CiviCRM {
	return &CiviCRM{
		client: client,
		config: config,
	}
}

func (c *CiviCRM) SearchContact(query SearchContactQuery) (response *SearchContactResponse, _ error) {
	response = &SearchContactResponse{}
	if req, err := c.buildQuery("Contact", "get", query); err != nil {
		return nil, err
	} else {
		return response, c.query(response, req)
	}
}

func (c *CiviCRM) ListContactMemberships(query ListMembershipsQuery) (response *ListMembershipsResponse, _ error) {
	response = &ListMembershipsResponse{}
	if req, err := c.buildQuery("Membership", "get", query); err != nil {
		return nil, err
	} else {
		return response, c.query(response, req)
	}
}

func (c *CiviCRM) CreateContact(contact string) {

}

func (c *CiviCRM) CreateContribution(contribution Contribution) (response *CreateContributionResponse, _ error) {
	response = &CreateContributionResponse{}
	if req, err := c.buildQuery("Contribution", "create", contribution); err != nil {
		return nil, err
	} else {
		return response, c.query(response, req)
	}
}

func (c *CiviCRM) buildQuery(entity string, action string, query interface{}) (*http.Request, error) {
	q := url.Values{}
	q.Add("entity", entity)
	q.Add("action", action)
	q.Add("api_key", c.config.UserKey)
	q.Add("key", c.config.SiteKey)
	if json, err := json.Marshal(query); err != nil {
		log.Println("Error marshalling query", err)
	} else {
		q.Add("json", string(json))
	}

	req, err := http.NewRequest("POST", c.config.URL, bytes.NewBufferString(q.Encode()))
	if err != nil {
		log.Println("Error building query", err)
		return req, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accepts", "application/json")

	return req, nil
}

func (c *CiviCRM) query(response Response, req *http.Request) error {
	if resp, err := c.client.Do(req); err != nil {
		log.Println("Error contacting CiviCRM", err)
		return err
	} else {
		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
			return err
		} else if response.Success() {
			return nil
		} else {
			return ResponseError{response.GetErrorMessage()}
		}
	}
}