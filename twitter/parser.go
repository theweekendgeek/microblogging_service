package twitter

import (
	"encoding/json"
)

type Parser struct{}

func (Parser) ParseResponse(response *[]byte, str interface{}) error {
	return json.Unmarshal(*response, &str)
}
