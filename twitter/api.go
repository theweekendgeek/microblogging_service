package twitter

import (
	. "doescher.ninja/twitter-service/config"
	"doescher.ninja/twitter-service/data"
	"doescher.ninja/twitter-service/utils"
	"fmt"
)

type ApiClient struct{}

func (ApiClient) RequestTweets(id string) *data.Tweets {
	url := fmt.Sprintf(Const().EndpointTimelineByID, id)

	timelineResponse := request[data.TimelineResponse](url)
	tweets := getTweets(timelineResponse)

	return &tweets
}

func (ApiClient) RequestUser(id string) *data.Profile {
	url := Const().EndpointUserByID + id

	userResponse := request[data.UserReponse](url)
	profile := getUser(userResponse)

	return &profile
}

func getUser(profileResponse data.UserReponse) data.Profile {
	return profileResponse.Data
}

func getTweets(timelineResponse data.TimelineResponse) data.Tweets {
	return timelineResponse.Tweets
}

func request[T any](url string) T {
	resByte := MakeRequest(url)

	var resObj T
	err := Parser{}.ParseResponse(resByte, &resObj)
	utils.FatalIfError(err)

	return resObj
}
