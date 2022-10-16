package twitter

import (
	c "doescher.ninja/twitter-service/config"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

//func GetData(endpoint string, response interface{}) {
func GetData() {
	ids, err := readUserIds()
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range ids {
		url := c.UserById + strconv.FormatInt(int64(v), 10)

		res, err := MakeRequest(url)
		if err != nil {
			log.Fatal(err)
		}

		var profileRes ProfileResponse
		err = json.Unmarshal(res, &profileRes)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(profileRes.Data.Id)
		fmt.Println(profileRes.Data.Username)
		fmt.Println(profileRes.Data.Name)
	}

}

func readUserIds() ([]int, error) {
	file, err := os.Open("users.json")
	if err != nil {
		log.Fatal(err)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var userIds []int
	err = json.Unmarshal(content, &userIds)
	return userIds, err
}

func getTweetsForUser() {
	// Get Id of last tweet saved
	// if none saved, then we just get the last 10 tweets
	// get the tweets
	// if pagination, then get he next page of tweets
	// if no pagination then finish
	// save tweets in database
	//

}
