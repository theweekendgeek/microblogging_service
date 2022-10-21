package persitence

import (
	. "doescher.ninja/twitter-service/config"
	"doescher.ninja/twitter-service/data"
)

func GetUserById(id string) (data.Profile, uint, error) {
	var user Profile

	err := getDb().Where(&Profile{
		TwitterId: id,
	}).First(&user).Error

	if err != nil {
		return data.Profile{}, user.ID, err
	}

	profile := matchProfile(user)
	return profile, user.ID, err

}

func GetLastUser() (uint, error) {
	var profile Profile
	err := getDb().Last(&profile).Error

	return profile.ID, err
}

func CreateProfile(profile *data.Profile) {
	modelProfile := Profile{Name: profile.Name, TwitterId: profile.Id, Username: profile.Username}

	result := getDb().Create(&modelProfile)
	FatalIfError(result.Error)
}

func CreateTweets(tweets *data.Tweets, userId uint) {
	var tweetModels []Tweet
	for _, v := range *tweets {
		tweetModels = append(tweetModels, matchTweetToModel(v, userId))
	}

	err := getDb().Create(&tweetModels).Error
	FatalIfError(err)
}

func matchTweetToModel(tweet data.Tweet, userid uint) Tweet {
	return Tweet{
		Text:      tweet.Text,
		ProfileID: userid,
	}
}

func matchProfile(model Profile) data.Profile {
	return data.Profile{
		Id:       model.TwitterId,
		Name:     model.Name,
		Username: model.Username,
	}

}

func DeleteTweets() {
	//goland:noinspection ALL
	getDb().Exec("DELETE FROM tweets")
}
