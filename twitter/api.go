package twitter

import (
	. "doescher.ninja/twitter-service/config"
	"doescher.ninja/twitter-service/data"
	"doescher.ninja/twitter-service/utils"
	"fmt"
)

func RequestTweets(id string) *data.Tweets {
	url := fmt.Sprintf(Const().EndpointTimelineByID, id)
	res := MakeRequest(url)

	var timelineResponse data.TimelineResponse

	err := Parser{}.ParseResponse(res, &timelineResponse)
	utils.FatalIfError(err)

	tweets := data.TimelineResponse{
		Tweets:   timelineResponse.Tweets,
		MetaData: timelineResponse.MetaData,
	}.Tweets
	return &tweets

}

func RequestProfile(id string) *data.Profile {
	res := MakeRequest(Const().EndpointUserByID + id)

	var profileResponse data.ProfileResponse
	err := Parser{}.ParseResponse(res, &profileResponse)
	utils.FatalIfError(err)

	profile := data.Profile{
		ID:       profileResponse.Data.ID,
		Name:     profileResponse.Data.Name,
		Username: profileResponse.Data.Username,
	}
	return &profile
}
