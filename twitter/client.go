package twitter

import (
	"io"
	"log"
	"net/http"
	"os"
)

type Profile struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type ProfileResponse struct {
	Data Profile `json:"data"`
}

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
