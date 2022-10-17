package twitter

import (
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
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+os.Getenv("BEARER"))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return io.ReadAll(resp.Body)
}
