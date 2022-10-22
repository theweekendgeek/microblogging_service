package business

import (
	"doescher.ninja/twitter-service/data"
	"doescher.ninja/twitter-service/persitence"
	"doescher.ninja/twitter-service/twitter"
	"doescher.ninja/twitter-service/utils"
	"sync"
)

func RequestAndSaveTweets() {
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
	if len(*tweets) > 0 {
		persitence.CreateTweets(tweets, userID)
	}

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

	params := twitter.QueryOptions{}

	tweet, notFoundError := persitence.GetLastSavedTweet(id)
	isNewUser := notFoundError != nil

	if !isNewUser {
		params.SinceID = tweet.TwitterID
	}

	var tweets data.Tweets
	for {
		//timeline := apiClient.RequestTweets(id, params)
		timeline := twitter.Request[data.TimelineResponse](twitter.GetUrlForId(id, params))

		if timeline.MetaData.ResultCount == 0 {
			break
		}

		tweets = append(tweets, timeline.Tweets...)

		if isNewUser || noFurtherTweets(timeline) {
			break
		}

		params.PaginationToken = timeline.MetaData.NextToken
	}

	return &tweets
}

func noFurtherTweets(timelinePointer *data.TimelineResponse) bool {
	return timelinePointer.MetaData.NextToken == ""
}

func getUser(id string) *data.Profile {

	return twitter.Request[data.Profile](twitter.BuildProfileUrl(id))
}
