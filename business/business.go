package business

import (
	"doescher.ninja/twitter-service/data"
	"doescher.ninja/twitter-service/persitence"
	"doescher.ninja/twitter-service/twitter"
	"doescher.ninja/twitter-service/utils"
	"sync"
)

var apiClient = twitter.APIClient{}

func RequestAndSaveTweets() {
	persitence.DeleteTweets()

	userIDs := utils.ReadUserIDs()

	wg := sync.WaitGroup{}
	wg.Add(len(userIDs))

	for _, id := range userIDs {
		go retrieveNewTweets(id, &wg)
	}

	wg.Wait()
}

func retrieveNewTweets(id string, wg *sync.WaitGroup) {
	_, userID, noRecordError := persitence.GetUserByID(id)
	if noRecordError != nil {
		userID = saveUser(id)
	}

	tweets := getTweetsForUser(id)
	persitence.CreateTweets(tweets, userID)

	wg.Done()
}

func saveUser(id string) uint {
	lastUserID, noRecordError := persitence.GetLastUser()

	var profileID uint
	if noRecordError != nil {
		profileID = 1
	} else {
		profileID = lastUserID + 1
	}

	profile := getUser(id)
	persitence.CreateUser(profile)
	return profileID
}

func getTweetsForUser(id string) *data.Tweets {

	// get latest tweet from database
	_, err := persitence.GetLastSavedTweet(id)
	utils.FatalIfError(err)
	// request new tweets since latest

	// paginate if necessary
	return apiClient.RequestTweets(id)

}

func getUser(id string) *data.Profile {
	return apiClient.RequestUser(id)
}
