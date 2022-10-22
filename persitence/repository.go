package persitence

import (
	"doescher.ninja/twitter-service/data"
	"doescher.ninja/twitter-service/utils"
)

// TODO: find a better way to return ids or models
func GetUserByID(twitterID string) (data.Profile, uint, error) {
	var user Profile

	err := getDb().Where(&Profile{
		TwitterId: twitterID,
	}).First(&user).Error

	if err != nil {
		return data.Profile{}, user.ID, err
	}

	return matchProfile(user), user.ID, err

}

func GetLastUser() (uint, error) {
	var profile Profile
	err := getDb().Last(&profile).Error

	return profile.ID, err
}

func CreateUser(profile *data.Profile) {
	modelProfile := Profile{Name: profile.Name, TwitterId: profile.ID, Username: profile.Username}

	result := getDb().Create(&modelProfile)
	utils.FatalIfError(result.Error)
}

func CreateTweets(tweets *data.Tweets, userID uint) {
	var tweetModels []Tweet
	for _, v := range *tweets {
		tweetModels = append(tweetModels, matchTweetToModel(v, userID))
	}

	err := getDb().Create(&tweetModels).Error
	utils.FatalIfError(err)
}

func GetLastSavedTweet(twitterId string) (data.Tweet, error) {
	var tweet Tweet

	_, modelId, err := GetUserByID(twitterId)

	err = getDb().Where(Tweet{ProfileID: modelId}).Last(&tweet).Error
	return matchModelToTweet(tweet), err
}

func matchModelToTweet(tweet Tweet) data.Tweet {
	return data.Tweet{
		TwitterID: tweet.TwitterID,
		Text:      tweet.Text,
	}
}

func matchTweetToModel(tweet data.Tweet, userid uint) Tweet {
	return Tweet{
		Text:      tweet.Text,
		TwitterID: tweet.TwitterID,
		ProfileID: userid,
	}
}

func matchProfile(model Profile) data.Profile {
	return data.Profile{
		ID:       model.TwitterId,
		Name:     model.Name,
		Username: model.Username,
	}

}

//func DeleteTweets() {
//	//goland:noinspection ALL
//	getDb().Exec("DELETE FROM tweets")
//}
