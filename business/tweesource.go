package business

import (
	"doescher.ninja/twitter-service/data"
	"doescher.ninja/twitter-service/persitence"
	"doescher.ninja/twitter-service/twitter"
)

type TweetSource struct{}

var apiClient = twitter.APIClient{}

func (ps TweetSource) GetUser(id string) *data.Profile {
	return apiClient.RequestUser(id)
}

func noFurtherTweets(timelinePointer *data.TimelineResponse) bool {
	return timelinePointer.MetaData.NextToken == ""
}

func getTweetsForNewUser(notFoundError error) bool {
	return notFoundError != nil
}

func (ps TweetSource) GetPostsForID(id string) *data.Posts {

	params := twitter.QueryOptions{}
	// get latest tweet from database
	tweet, notFoundError := persitence.GetLastSavedTweet(id)

	// if we have tweets, set a SinceId to just get new ones
	if notFoundError == nil {
		params.SinceID = tweet.PostID
	}

	var tweets data.Posts

	loop := 0
	for {
		// request new tweets since latest
		timelinePointer := apiClient.RequestTweets(id, params)
		if timelinePointer.MetaData.ResultCount == 0 {
			break
		}
		tweets = append(tweets, timelinePointer.Posts...)

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
