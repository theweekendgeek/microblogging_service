package twitter

type Profile struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type Data struct {
}

type ProfileResponse struct {
	Data Profile `json:"data"`
}

type TimelineResponse struct {
}
