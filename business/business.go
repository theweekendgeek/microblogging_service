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
	userIds, err := config.ReadUserIds()
	config.FatalIfError(err)

	wg := sync.WaitGroup{}
	wg.Add(len(userIds))

	for _, id := range userIds {
		go retrieveNewTweets(id)

		wg.Done()
	}
}

func retrieveNewTweets(id string) {
	_, profileId, noRecordError := persitence.GetUserById(id)
	if noRecordError != nil {
		profileId = createProfile(id)
	}

	tweets := GetTweetsForUser(id)
	persitence.CreateTweets(tweets, profileId)
}

func createProfile(id string) uint {
	lastUserId, noRecordError := persitence.GetLastUser()

	var profileId uint
	if noRecordError != nil {
		profileId = 1
	} else {
		profileId = lastUserId + 1
	}

	profile := GetUserProfile(id)
	persitence.CreateProfile(profile)
	return profileId
}

func GetTweetsForUser(id string) *data.Tweets {
	return twitter.RequestTweets(id)

}

func GetUserProfile(id string) *data.Profile {
	return twitter.RequestProfile(id)
}
