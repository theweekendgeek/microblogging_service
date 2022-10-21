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
		noRecordError, _, profileId := persitence.GetUserById(id)
		if noRecordError != nil {
			profileId = createProfile(profileId, id)
		}

		tweets := GetTweetsForUser(id)
		go persitence.CreateTweets(&tweets, profileId)
		wg.Done()
	}
}

func createProfile(profileId uint, id string) uint {
	noRecordError, lastUserId := persitence.GetLastUser()
	if noRecordError != nil {
		profileId = 1
	} else {
		profileId = lastUserId + 1
	}

	profile := GetUserProfile(id)
	persitence.CreateProfile(&profile)
	return profileId
}

func GetTweetsForUser(id string) data.Tweets {
	return twitter.RequestTweets(id)

}

func GetUserProfile(id string) data.Profile {
	return twitter.RequestProfile(id)
}
