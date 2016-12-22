package main

import (
	"net/http"
	"encoding/json"
	"os"
	"log"
	"net/url"
	"strconv"
)

type searchResponse struct {
	IsError int `json:"is_error" binding:"required"`
	Version int `json:"version" binding:"required"`
	Count int `json:"count" binding:"required"`
	Id int `json:"id"`
}

type listMembershipsResponse struct {
	IsError int `json:"is_error" binding:"required"`
	Version int `json:"version" binding:"required"`
	Count int `json:"count" binding:"required"`
	Values map[int]membership `json:"values"`
}

type membership struct {
	Id int `json:"id,string" binding:"required"`
	ContactId int `json:"contact_id,string" binding:"required"`
	MembershipTypeId int `json:"membership_type_id,string" binding:"required"`
	JoinDate string `json:"join_date" binding:"required"`
	StartDate string `json:"start_date" binding:"required"`
	EndDate string `json:"end_date" binding:"required"`
	StatusId int `json:"status_id,string" binding:"required"`
	IsOverride int `json:"is_override,string"`
}

func searchContact(email string) (searchResponse, error) {
	var response searchResponse
	if req, err := buildSearchContactQuery(email); err != nil {
		return response, err
	} else {
		return response, query(&response, req)
	}
}

func listContactMemberships(contactId int) (listMembershipsResponse, error) {
	var response listMembershipsResponse
	if req, err := buildListContactMembershipsQuery(contactId); err != nil {
		return response, err
	} else {
		return response, query(&response, req)
	}
}


func createContact(contact string) {

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

type Customizer func(q *url.Values)

func buildBaseQuery(entity string, action string, customizer Customizer) (*http.Request, error) {
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
	log.Println(req)
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