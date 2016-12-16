package main

import (
	"net/http"
	"encoding/json"
)

const url = "https://crm.wikimedia.fr"

type searchResponse struct {

}

func SearchContact(email string) {
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var response searchResponse
	err = decoder.Decode(&response)
	if err != nil {
		return err
	}
	return response
}

func ListContactMemberships(contact string) {

}

func CreateContact(contact string) {

}
