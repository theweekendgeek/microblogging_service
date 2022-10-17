package twitter

import (
	c "doescher.ninja/twitter-service/config"
	"doescher.ninja/twitter-service/orm"
	"encoding/json"
	"gorm.io/gorm"
	"io"
	"log"
	"os"
	"strconv"
)

func GetProfiles() {
	ids, err := readUserIds()
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range ids {
		profileResponse, err := RequestData(v)
		if err != nil {
			log.Fatal(err)
		}

		a, ok := profileResponse.(ProfileResponse)
		if !ok {
			log.Fatal(ok)
		}

		profileModel := orm.GormProfile{
			Model:    gorm.Model{},
			Id:       a.Data.Id,
			Username: a.Data.Username,
			Name:     a.Data.Name,
		}

		_ = orm.GetDb().Create(&profileModel)
	}
}

//func RequestData(endpoint string, response interface{}) {
func RequestData(v int) (interface{}, error) {

	url := c.UserById + strconv.FormatInt(int64(v), 10)
	res, err := MakeRequest(url)
	if err != nil {
		log.Fatal(err)
	}

	var profileResponse ProfileResponse
	err = Parser{}.ParseResponse(res, &profileResponse)

	return profileResponse, err
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
