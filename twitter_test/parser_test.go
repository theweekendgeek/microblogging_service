package twitter_test

import (
	"doescher.ninja/twitter-service/data"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)
import "doescher.ninja/twitter-service/twitter"

var parser twitter.Parser

func init() {
	parser = twitter.Parser{}
}

func TestParseProfileResponse(t *testing.T) {
	var profileResponse data.ProfileResponse
	response := []byte(`
		{
			"data": {
				"id": "11348282",
				"name": "NASA",
				"username": "NASA"
			}
		}
	`)

	err := parser.ParseResponse(response, &profileResponse)
	if err != nil {
		t.Fatal("error")
	}

	fmt.Println(profileResponse)
	if assert.NotNil(t, profileResponse.Data) {
		assert.Equal(t, "NASA", profileResponse.Data.Name)
		assert.Equal(t, "11348282", profileResponse.Data.Id)
		assert.Equal(t, "NASA", profileResponse.Data.Username)

	}

	//if twitter.ProfileResponse(parsedResponse).Data.Id == "" {
	//
	//}

}

func TestParsesTimelines(t *testing.T) {
	response := []byte(`{
	    "data": [
	        {
	            "edit_history_tweet_ids": [
	                "1582051114941566976"
	            ],
	            "id": "1582051114941566976",
	            "text": "Your delivery is being prepared. 📦\n\nNorthrop Grumman's next cargo mission to the @Space_Station will deliver new experiments studying the dynamics of mudflows, growing crops in space, and fertility treatments—all for the benefit of humanity: https://t.co/bGCuPQv6a1 https://t.co/JLZWuwCJJx"
	        },
	        {
	            "edit_history_tweet_ids": [
	                "1581628790878281729"
	            ],
	            "id": "1581628790878281729",
	            "text": "RT @NASASolarSystem: ✅NASA’s #LucyMission completed its 1st Earth Gravity Assist, passing just 220 miles/350 km above Earth’s surface!\n\nIn…"
	        },
	        {
	            "edit_history_tweet_ids": [
	                "1581389969154658304"
	            ],
	            "id": "1581389969154658304",
	            "text": "Our workforce is our strength. Check out our Twitter Moment for a recap of #HispanicHeritageMonth and #MesDeLaHerenciaHispana 2022, and follow @NASA_es for Spanish-language coverage all year round. https://t.co/YsLuM7u5vJ"
	        },
	        {
	            "edit_history_tweet_ids": [
	                "1581320295318622209"
	            ],
	            "id": "1581320295318622209",
	            "text": "“I try to help people see that there’s an opportunity, and that if you work hard and are willing to evoke whatever it takes to get there, you can get it.”\nAndres Rivera is a systems engineer for the @EuropaClipper mission. Celebrate #HispanicHeritageMonth: https://t.co/6jZCsmwUD5 https://t.co/NDklOV5wLh"
	        },
	        {
	            "edit_history_tweet_ids": [
	                "1581302372453605377"
	            ],
	            "id": "1581302372453605377",
	            "text": "\"The last time I was home, I found a little note that I had written when I was a 13-year-old kid. It said, 'My dream job in the future is to be a programmer.'\"\nDr. Ivan Perez Dominguez is a research scientist at @NASAAmes. Celebrate #HispanicHeritageMonth: https://t.co/d58POvHyd8 https://t.co/EVVt21xYvg"
	        },
	        {
	            "edit_history_tweet_ids": [
	                "1581068720092631040"
	            ],
	            "id": "1581068720092631040",
	            "text": "This Week at NASA:\n\n• Confirmed! DART’s impact changed the orbit of asteroid Dimorphos\n• #Crew4 heads to Earth after 170 days in space\n• @NASAWebb sees cosmic \"tree rings\"\n• #Artemis I sets a new launch date\n\nWant more space in your life? Subscribe: https://t.co/MyG37QAe7m https://t.co/tQsM528BP3"
	        },
	        {
	            "edit_history_tweet_ids": [
	                "1581056289282482176"
	            ],
	            "id": "1581056289282482176",
	            "text": "RT @nasahqphoto: Welcome home! Crew-4 NASA astronauts @astro_watkins, @astro_farmerbob, @astro_kjell, and @ESA’s @AstroSamantha are seen in…"
	        },
	        {
	            "edit_history_tweet_ids": [
	                "1581050294208581633"
	            ],
	            "id": "1581050294208581633",
	            "text": "LIVE: Tune in for an update from NASA and @SpaceX leaders reviewing today's #Crew4 splashdown off the coast of Florida: https://t.co/IRdNh3Bqb7"
	        },
	        {
	            "edit_history_tweet_ids": [
	                "1581046644467847168"
	            ],
	            "id": "1581046644467847168",
	            "text": "75 years ago, NACA—NASA’s predecessor—broke the sound barrier in our first experimental aircraft, the X-1. Now, via @NASAAero’s Quesst mission, we aim to do it again, but much quieter, to bring a return to commercial supersonic flight over land: https://t.co/LOWWTbFpSu https://t.co/Ezoi6Qboib"
	        },
	        {
	            "edit_history_tweet_ids": [
	                "1581043668206891008"
	            ],
	            "id": "1581043668206891008",
	            "text": "After 170 days in orbit, our SpaceX Crew-4 astronauts safely splashed down at 4:55pm ET (20:55 UTC) Friday off the coast of Jacksonville, Florida, completing our fourth @Commercial_Crew mission to the International @Space_Station: https://t.co/XLOeKCjrTM https://t.co/yH5KYONXZa"
	        }
	    ],
	    "meta": {
	        "next_token": "7140dibdnow9c7btw423wugb5ysor0cpuq8gmvxycpcmw",
	        "result_count": 10,
	        "newest_id": "1582051114941566976",
	        "oldest_id": "1581043668206891008"
	    }
	 }`)
	var timelineResponse = data.TimelineResponse{}

	_ = parser.ParseResponse(response, &timelineResponse)

	if assert.NotNil(t, timelineResponse.Tweets) && assert.NotNil(t, timelineResponse.MetaData) {
		fmt.Println(timelineResponse.Tweets[0].Text)
		assert.Equal(t, "1582051114941566976", timelineResponse.Tweets[0].Id)
		assert.Equal(t, "Your delivery is being prepared. 📦\n\nNorthrop Grumman's next cargo mission to the @Space_Station will deliver new experiments studying the dynamics of mudflows, growing crops in space, and fertility treatments—all for the benefit of humanity: https://t.co/bGCuPQv6a1 https://t.co/JLZWuwCJJx", timelineResponse.Tweets[0].Text)
		assert.Equal(t, "1581056289282482176", timelineResponse.Tweets[6].Id)
		assert.Equal(t, "RT @nasahqphoto: Welcome home! Crew-4 NASA astronauts @astro_watkins, @astro_farmerbob, @astro_kjell, and @ESA’s @AstroSamantha are seen in…", timelineResponse.Tweets[6].Text)
		assert.Equal(t, "7140dibdnow9c7btw423wugb5ysor0cpuq8gmvxycpcmw", timelineResponse.MetaData.NextToken)
		assert.Equal(t, 10, timelineResponse.MetaData.ResultCount)
	}

}
