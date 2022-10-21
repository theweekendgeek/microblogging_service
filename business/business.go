package business

import (
	"doescher.ninja/twitter-service/config"
	"doescher.ninja/twitter-service/data"
	"doescher.ninja/twitter-service/persitence"
	"doescher.ninja/twitter-service/twitter"
	"sync"
)

func GetNewTweets() {
	persitence.DeleteTweets()
	userIDs, err := config.ReadUserIDs()
	config.FatalIfError(err)

	wg := sync.WaitGroup{}
	wg.Add(len(userIDs))

	for _, id := range userIDs {
		go retrieveNewTweets(id, &wg)

	}
	wg.Wait()
}

func retrieveNewTweets(id string, wg *sync.WaitGroup) {
	_, profileID, noRecordError := persitence.GetUserByID(id)
	if noRecordError != nil {
		profileID = createProfile(id)
	}

	tweets := GetTweetsForUser(id)
	persitence.CreateTweets(tweets, profileID)

	wg.Done()
}

func createProfile(id string) uint {
	lastUserID, noRecordError := persitence.GetLastUser()

	var profileID uint
	if noRecordError != nil {
		profileID = 1
	} else {
		profileID = lastUserID + 1
	}

	profile := GetUserProfile(id)
	persitence.CreateProfile(profile)
	return profileID
}

func GetTweetsForUser(id string) *data.Tweets {
	return twitter.RequestTweets(id)

}

func GetUserProfile(id string) *data.Profile {
	return twitter.RequestProfile(id)
}
