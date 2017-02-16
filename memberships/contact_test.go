package memberships

import (
	"testing"
	"github.com/wikimedia-france/wmfr-adhesions/civicrm"
	"github.com/wikimedia-france/wmfr-adhesions/iraiser"
	"github.com/stretchr/testify/assert"
)

func TestGetContactWithCrmError(t *testing.T) {
	contactGetter = func(*civicrm.GetContactQuery) (*civicrm.GetContactResponse, error) {
		response := &civicrm.GetContactResponse{}
		response.IsError = 1
		return response, assert.AnError
	}
	donation := iraiser.Donation{
	}
	matches, err := getContact(&donation)
	assert.Equal(t, 0, matches)
	assert.Error(t, err)
}

func TestGetContactNotFound(t *testing.T) {
	contactGetter = func(*civicrm.GetContactQuery) (*civicrm.GetContactResponse, error) {
		response := &civicrm.GetContactResponse{}
		return response, nil
	}
	donation := iraiser.Donation{
	}
	matches, err := getContact(&donation)
	assert.Equal(t, 0, matches)
	assert.Error(t, err)
}

func TestGetContactSeveralResponses(t *testing.T) {
	contactGetter = func(*civicrm.GetContactQuery) (*civicrm.GetContactResponse, error) {
		response := &civicrm.GetContactResponse{}
		response.Count = 2
		return response, nil
	}
	donation := iraiser.Donation{
	}
	matches, err := getContact(&donation)
	assert.Equal(t, 0, matches)
	assert.Error(t, err)
}

func TestGetContactExactlyOne(t *testing.T) {
	contactGetter = func(*civicrm.GetContactQuery) (*civicrm.GetContactResponse, error) {
		response := &civicrm.GetContactResponse{}
		response.Count = 1
		response.Id = 42
		return response, nil
	}
	donation := iraiser.Donation{
	}
	matches, err := getContact(&donation)
	assert.Equal(t, 42, matches)
	assert.NoError(t, err)
}
