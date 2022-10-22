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
	tweets := mapTimeline(timelineResponse)

	return &tweets
}

func (ApiClient) RequestUser(id string) *data.Profile {
	url := Const().EndpointUserByID + id

	profileResponse := request[data.ProfileResponse](url)
	profile := mapProfile(profileResponse)

	return &profile
}

func mapProfile(profileResponse data.ProfileResponse) data.Profile {
	return data.Profile{
		ID:       profileResponse.Data.ID,
		Name:     profileResponse.Data.Name,
		Username: profileResponse.Data.Username,
	}
}

func mapTimeline(timelineResponse data.TimelineResponse) data.Tweets {
	return data.TimelineResponse{
		Tweets:   timelineResponse.Tweets,
		MetaData: timelineResponse.MetaData,
	}.Tweets
}

func request[T any](url string) T {
	resByte := MakeRequest(url)

	var resObj T
	err := Parser{}.ParseResponse(resByte, &resObj)
	utils.FatalIfError(err)

	return resObj
}
