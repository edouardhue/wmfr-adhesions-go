package civicrm

import (
	"net/http"
	"encoding/json"
	"log"
	"net/url"
	"bytes"
	"net/http/httputil"
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

func (c *CiviCRM) buildQuery(entity string, action string, query interface{}) (*http.Request, error) {
	q := url.Values{}
	q.Add("entity", entity)
	q.Add("action", action)
	q.Add("api_key", c.config.UserKey)
	q.Add("key", c.config.SiteKey)
	if jsonQuery, err := json.Marshal(query); err != nil {
		log.Println("Error marshalling query", err)
	} else {
		q.Add("json", string(jsonQuery))
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

func (c *CiviCRM) query(response Status, req *http.Request) error {
	dump, _ := httputil.DumpRequestOut(req, true)
	resp, err := c.client.Do(req)
	if err != nil {
		log.Println("Error contacting CiviCRM", err)
		return err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return err
	}
	if response.Success() {
		return nil
	}
	unescaped, _ := url.QueryUnescape(string(dump))
	return ResponseError{
		Request: unescaped,
		Message: response.GetErrorMessage(),
	}
}