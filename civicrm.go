package main

import (
	"net/http"
	"encoding/json"
	"os"
	"log"
)

type SearchResponse struct {
	IsError bool `json:"isError" binding:"required"`
	Version int `json:"version" binding:"required"`
	Count int `json:"count" binding:"required"`
	Id int `json:"id"`
}

func buildSearchContactQuery(email string) (*http.Request, error) {
	req, err := http.NewRequest("GET", os.Getenv("CIVI_URL"), nil)
	if err != nil {
		return req, err
	}
	q := req.URL.Query()
	q.Add("entity", "Contact")
	q.Add("action", "get")
	q.Add("json", "1")
	q.Add("api_key", os.Getenv("CIVI_API_KEY"))
	q.Add("key", os.Getenv("CIVI_KEY"))
	q.Add("email", email)
	req.URL.RawQuery = q.Encode()
	return req, nil
}

func SearchContact(email string) (SearchResponse, error) {
	if req, err := buildSearchContactQuery(email); err != nil {
		log.Println("Error building query", err)
		return SearchResponse{}, err
	} else {
		if resp, err := http.DefaultClient.Do(req); err != nil {
			log.Println("Error contacting CiviCRM", err)
			return SearchResponse{}, err
		} else {
			defer resp.Body.Close()
			decoder := json.NewDecoder(resp.Body)
			var response SearchResponse
			if err := decoder.Decode(&response); err != nil {
				log.Println("Could not parse CiviCRM response", err)
				return SearchResponse{}, err
			} else {
				return response, nil
			}
		}
	}
}

func ListContactMemberships(contact string) {

}

func CreateContact(contact string) {

}
