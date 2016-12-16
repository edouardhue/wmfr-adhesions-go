package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"os"
)

const url = "https://crm.wikimedia.fr"

type searchResponse struct {

}

func SearchContact(email string) {
	if resp, err := http.DefaultClient.Get(url); err != nil {
		fmt.Fprintln(os.Stderr, "Error contacting CiviCRM", err)
	} else {
		defer resp.Body.Close()
		decoder := json.NewDecoder(resp.Body)
		var response searchResponse
		if err := decoder.Decode(&response); err != nil {
			fmt.Fprintln(os.Stderr, "Could not parse CiviCRM response", err)
		} else {
			return response
		}
	}
}

func ListContactMemberships(contact string) {

}

func CreateContact(contact string) {

}
