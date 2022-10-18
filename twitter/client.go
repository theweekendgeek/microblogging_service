package twitter

import (
	. "doescher.ninja/twitter-service/config"
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

func MakeRequest(url string) ([]byte, error) {

	req, err := http.NewRequest("GET", url, nil)
	FatalIfError(err)

	req.Header.Add("Authorization", "Bearer "+os.Getenv("BEARER"))
	resp, err := client.Do(req)
	FatalIfError(err)

	if resp.StatusCode != 200 {
		log.Fatal(fmt.Sprintf("Got an Error with StatusCode %d", resp.StatusCode))
	}

	return io.ReadAll(resp.Body)
}
