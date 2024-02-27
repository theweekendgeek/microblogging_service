package business

import "doescher.ninja/twitter-service/data"

type MastodonSource struct{}

func (m MastodonSource) GetUser(id string) *data.Profile {
	//TODO implement me
	panic("implement me")
}

func (m MastodonSource) GetPostsForID(id string) *data.Posts {
	//TODO implement me
	panic("implement me")
}
