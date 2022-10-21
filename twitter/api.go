package twitter

import (
	. "doescher.ninja/twitter-service/config"
	"doescher.ninja/twitter-service/data"
	"fmt"
)

func RequestTweets(id string) *data.Tweets {
	url := fmt.Sprintf(Const().TimelineByID, id)
	res := MakeRequest(url)

	var timelineResponse data.TimelineResponse

	err := Parser{}.ParseResponse(res, &timelineResponse)
	FatalIfError(err)

	tweets := data.TimelineResponse{
		Tweets:   timelineResponse.Tweets,
		MetaData: timelineResponse.MetaData,
	}.Tweets
	return &tweets

}

func RequestProfile(id string) *data.Profile {
	res := MakeRequest(Const().UserByID + id)

	var profileResponse data.ProfileResponse
	err := Parser{}.ParseResponse(res, &profileResponse)
	FatalIfError(err)

	profile := data.Profile{
		ID:       profileResponse.Data.ID,
		Name:     profileResponse.Data.Name,
		Username: profileResponse.Data.Username,
	}
	return &profile
}
