// Package data holds the structs to handle the api responses
package data

// Profile represents a Twitter users profile
type Profile struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

// Post represents a single tweet
type Post struct {
	PostID      string `json:"id"`
	PostContent string `json:"text"`
}

// Meta holds meta-information about a user's timeline
type Meta struct {
	NextToken   string `json:"next_token"`
	ResultCount int    `json:"result_count"`
}

// UserReponse is the result of requesting a user's profile
type UserReponse struct {
	Data Profile `json:"data"`
}

// Posts is a slice of post
type Posts []Post

// TimelineResponse is the result of requesting a user's timeline
type TimelineResponse struct {
	Posts    Posts `json:"data"`
	MetaData Meta  `json:"meta"`
}
