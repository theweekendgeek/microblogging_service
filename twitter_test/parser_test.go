package twitter_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)
import "doescher.ninja/twitter-service/twitter"

func TestParseProfileResponse(t *testing.T) {
	type Timeline struct{}
	parser := twitter.Parser{}
	var profileResponse twitter.ProfileResponse
	response := []byte(`
		{
			"data": {
				"id": "11348282",
				"name": "NASA",
				"username": "NASA"
			}
		}
	`)

	err := parser.ParseProfile(response, &profileResponse)
	if err != nil {
		t.Fatal("error")
	}

	fmt.Println(profileResponse)
	if assert.NotNil(t, profileResponse.Data) {
		assert.Equal(t, "NASA", profileResponse.Data.Name)
		assert.Equal(t, "11348282", profileResponse.Data.Id)
		assert.Equal(t, "NASA", profileResponse.Data.Username)

	}

	//if twitter.ProfileResponse(parsedResponse).Data.Id == "" {
	//
	//}

}

func TestParsesProfiles(t *testing.T) {

}
