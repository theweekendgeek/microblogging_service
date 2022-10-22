package twitter

import (
	"doescher.ninja/twitter-service/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var client *http.Client

func init() {
	client = &http.Client{}
}

func MakeRequest(url string) *[]byte {
	req, err := http.NewRequest("GET", url, nil)
	utils.FatalIfError(err)

	req.Header.Add("Authorization", "Bearer "+os.Getenv("BEARER"))
	resp, err := client.Do(req)
	utils.FatalIfError(err)

	if resp.StatusCode != 200 {
		log.Fatal(fmt.Sprintf("Got an Error with StatusCode %d", resp.StatusCode))
	}

	a, err := io.ReadAll(resp.Body)
	utils.FatalIfError(err)
	return &a
}
