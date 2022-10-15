package twitter

import (
	c "doescher.ninja/twitter-service/config"
	"encoding/json"
	"fmt"
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

func Request() {
	client := &http.Client{}

	username := "BarackObama"
	req, err := http.NewRequest("GET", c.UserByName+username, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+os.Getenv("BEARER"))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Body)
	//var a []byte
	//fmt.Println(resp.Body.Read(a))

	//fmt.Println(a)
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resBody)

	var profileRes ProfileResponse

	err = json.Unmarshal(resBody, &profileRes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(profileRes.Data.Id)
	fmt.Println(profileRes.Data.Username)
	fmt.Println(profileRes.Data.Name)
}
