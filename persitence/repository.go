package persitence

import (
	"doescher.ninja/twitter-service/data"
	"doescher.ninja/twitter-service/utils"
)

// TODO: find a better way to return ids or models
func GetUserByID(twitterID string) (data.Profile, uint, error) {
	// TODO: find a better way to return ids or models
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

// TODO: breaks on first run for a user
func CreatePosts(tweets *data.Posts, userID uint) {
	var tweetModels []Tweet
	for _, v := range *tweets {
		tweetModels = append(tweetModels, matchTweetToModel(v, userID))
	}

	err := getDb().Create(&tweetModels).Error
	utils.FatalIfError(err)
}

// GetLastSavedTweet get the newest tweet for a user
func GetLastSavedTweet(twitterID string) (data.Post, error) {
	var tweet Tweet

	_, modelID, err := GetUserByID(twitterID)
	utils.FatalIfError(err)

	err = getDb().Where(Tweet{ProfileID: modelID}).Order("twitter_id DESC").Take(&tweet).Error
	return matchModelToTweet(tweet), err
}

func matchModelToTweet(tweet Tweet) data.Post {
	return data.Post{
		PostID:      tweet.TwitterID,
		PostContent: tweet.Text,
	}
}

func matchTweetToModel(tweet data.Post, userid uint) Tweet {
	return Tweet{
		Text:      tweet.PostContent,
		TwitterID: tweet.PostID,
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
