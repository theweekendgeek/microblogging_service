package twitter

import (
	. "doescher.ninja/twitter-service/config"
	"doescher.ninja/twitter-service/data"
	"doescher.ninja/twitter-service/utils"
	"fmt"
)

type APIClient struct{}

func (APIClient) RequestTweets(id string, opts QueryOptions) *data.TimelineResponse {
	url := buildTimelineUrl(opts)
	url = fmt.Sprintf(url, id)

	timelineResponse := request[data.TimelineResponse](url)
	tweets := getTweets(timelineResponse)

	return &tweets
}

func (APIClient) RequestUser(id string) *data.Profile {
	url := Const().EndpointUserByID + id

	userResponse := request[data.UserReponse](url)
	profile := getUser(userResponse)

	return &profile
}

func getUser(profileResponse data.UserReponse) data.Profile {
	return profileResponse.Data
}

func getTweets(timelineResponse data.TimelineResponse) data.TimelineResponse {
	return timelineResponse
}

func request[T any](url string) T {
	resByte := MakeRequest(url)

	var resObj T
	err := Parser{}.ParseResponse(resByte, &resObj)
	utils.FatalIfError(err)

	return resObj
}
