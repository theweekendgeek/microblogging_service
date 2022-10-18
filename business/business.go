package business

import (
	"doescher.ninja/twitter-service/config"
	"doescher.ninja/twitter-service/data"
	"doescher.ninja/twitter-service/persitence"
	"doescher.ninja/twitter-service/twitter"
)

func GetNewTweets() {
	persitence.DeleteTweets()
	userIds, err := config.ReadUserIds()
	config.FatalIfError(err)

	for _, id := range userIds {
		noRecordError, _, profileId := persitence.GetUserById(id)
		if noRecordError != nil {
			profileId = createProfile(profileId, id)
		}

		tweets := GetTweetsForUser(id)
		persitence.CreateTweets(tweets, profileId)
	}
}

func createProfile(profileId uint, id string) uint {
	noRecordError, lastUserId := persitence.GetLastUser()
	if noRecordError != nil {
		profileId = 1
	} else {
		profileId = lastUserId + 1
	}

	persitence.CreateProfile(GetUserProfile(id))
	return profileId
}

func GetTweetsForUser(id string) data.Tweets {
	return twitter.RequestTweets(id)

}

func GetUserProfile(id string) data.Profile {
	return twitter.RequestProfile(id)
}
