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
	//persitence.DeleteTweets()

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
	// get latest tweet from database
	tweet, notFoundError := persitence.GetLastSavedTweet(id)

	// if we have tweets, set a SinceId to just get new ones
	if notFoundError == nil {
		params.SinceID = tweet.TwitterID
	}

	var tweets data.Tweets

	loop := 0
	for {
		// request new tweets since latest
		timelinePointer := apiClient.RequestTweets(id, params)
		tweets = append(tweets, timelinePointer.Tweets...)

		if getTweetsForNewUser(notFoundError) || noFurtherTweets(timelinePointer) {
			break
		}

		// paginate if necessary
		params.PaginationToken = timelinePointer.MetaData.NextToken
		loop++

		if loop == 4 { // limit number of loops to 5 for now
			break
		}
	}

	return &tweets

}

func noFurtherTweets(timelinePointer *data.TimelineResponse) bool {
	return timelinePointer.MetaData.NextToken == ""
}

func getTweetsForNewUser(notFoundError error) bool {
	return notFoundError != nil
}

func getUser(id string) *data.Profile {
	return apiClient.RequestUser(id)
}
