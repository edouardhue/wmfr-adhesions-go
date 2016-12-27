package civicrm

import (
	"net/http"
	"encoding/json"
	"os"
	"log"
	"net/url"
	"strconv"
)

func SearchContact(email string) (SearchResponse, error) {
	var response SearchResponse
	if req, err := buildSearchContactQuery(email); err != nil {
		return response, err
	} else {
		return response, query(&response, req)
	}
}

func ListContactMemberships(contactId int) (ListMembershipsResponse, error) {
	var response ListMembershipsResponse
	if req, err := buildListContactMembershipsQuery(contactId); err != nil {
		return response, err
	} else {
		return response, query(&response, req)
	}
}


func CreateContact(contact string) {

}

func buildSearchContactQuery(email string) (*http.Request, error) {
	return buildBaseQuery("Contact", "get", func(q *url.Values) {
		q.Add("email", email)
	})
}

func buildListContactMembershipsQuery(contactId int) (*http.Request, error) {
	return buildBaseQuery("Membership", "get", func(q *url.Values) {
		q.Add("contact_id", strconv.Itoa(contactId))
	})
}

type customizer func(q *url.Values)

func buildBaseQuery(entity string, action string, customizer customizer) (*http.Request, error) {
	req, err := http.NewRequest("GET", os.Getenv("CIVI_URL"), nil)
	if err != nil {
		log.Println("Error building query", err)
		return req, err
	}
	q := req.URL.Query()
	q.Add("entity", entity)
	q.Add("action", action)
	q.Add("json", "1")
	q.Add("api_key", os.Getenv("CIVI_API_KEY"))
	q.Add("key", os.Getenv("CIVI_KEY"))
	customizer(&q)
	req.URL.RawQuery = q.Encode()
	return req, nil
}

func query(response interface{}, req *http.Request) error {
	if resp, err := http.DefaultClient.Do(req); err != nil {
		log.Println("Error contacting CiviCRM", err)
		return err
	} else {
		defer resp.Body.Close()
		return json.NewDecoder(resp.Body).Decode(&response)
	}
}