package iraiser

import (
	"testing"
	"encoding/json"
	"io/ioutil"
)

func TestDonationJsonBinding(t *testing.T) {
	if bytes, err := ioutil.ReadFile("sample_donation.json"); err != nil {
		t.Fatal(err)
	} else {
		donation := Donation{}
		if err := json.Unmarshal(bytes, &donation); err != nil {
			t.Fatal(err)
		}
	}
}