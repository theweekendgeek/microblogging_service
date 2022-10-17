package twitter

import (
	"encoding/json"
	"fmt"
)

type Parser struct{}

func (Parser) ParseResponse(response []byte, str interface{}) error {
	return json.Unmarshal(response, &str)
}

func (p Parser) ParseProfile(data []byte, profile *ProfileResponse) error {
	fmt.Println(data, profile)
	return p.ParseResponse(data, profile)
}

func (p Parser) ParseTimeline(data []byte, timeline *TimelineResponse) error {
	return p.ParseResponse(data, timeline)
}
