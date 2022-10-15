package twitter

import (
	c "doescher.ninja/twitter-service/config"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
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
	ids := readUserIds()

	client := &http.Client{}

	//username := "BarackObama"
	for _, v := range ids {
		//fmt.Println(string(v))
		//fmt.Println()
		resustIds(strconv.FormatInt(int64(v), 10), client)
	}

	//resustIds(username, client)
}

func resustIds(username string, client *http.Client) {
	fmt.Println(c.UserById + username)
	//os.Exit(0)
	req, err := http.NewRequest("GET", c.UserById+username, nil)
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

func readUserIds() []int {
	contents, err := os.Open("users.json")
	if err != nil {
		log.Fatal(err)
	}

	t, err := io.ReadAll(contents)
	if err != nil {
		log.Fatal(err)
	}

	var userIds []int
	err = json.Unmarshal(t, &userIds)
	if err != nil {
		log.Fatal(err)
	}

	return userIds
}
