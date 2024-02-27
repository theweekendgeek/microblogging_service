package business

import (
	"doescher.ninja/twitter-service/data"
	"doescher.ninja/twitter-service/persitence"
	"doescher.ninja/twitter-service/utils"
	"sync"
)

type OnlineSource interface {
	GetPostsForID(id string) *data.Posts
	GetUser(id string) *data.Profile
}

func RequestAndSaveTweets() {
	b := utils.DataSource{}
	userIDs := b.GetProfileIDs()

	wg := sync.WaitGroup{}
	wg.Add(len(userIDs))

	postSource := TweetSource{}
	for _, id := range userIDs {
		go retrieveNewPosts(id, &wg, postSource)
	}

	wg.Wait()
}

func GetUserIDs(a utils.DataSourceInterface) []string {
	return a.GetProfileIDs()
}

func retrieveNewPosts(id string, wg *sync.WaitGroup, ps OnlineSource) {
	_, userID, noRecordError := persitence.GetUserByID(id)
	if noRecordError != nil {
		userID = saveUser(id, ps)
	}

	posts := ps.GetPostsForID(id)
	if len(*posts) > 0 {
		persitence.CreatePosts(posts, userID)
	}

	wg.Done()
}

func saveUser(id string, ps OnlineSource) uint {
	lastUserID, noRecordError := persitence.GetLastUser()

	var profileID uint
	if noRecordError != nil {
		profileID = 1
	} else {
		profileID = lastUserID + 1
	}

	profile := ps.GetUser(id)
	persitence.CreateUser(profile)
	return profileID
}
