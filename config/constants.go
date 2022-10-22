package config

var baseURL = "https://api.twitter.com"
var version = "/2"

var getUserByID = "/users/"
var timeLineByID = "/users/%s/tweets" //+ fmt.Sprintf("%d", Conf().MaxTweets)

type Constants struct {
	EndpointUserByID     string // uri to get a user profile by id
	EndpointTimelineByID string // uri to get a user's timeline by id
	UsersFileLocal       string // path to the user id file on local machine
	UsersFileCloud       string // path to the user id file on gcp cloud function
	EnvFileProd          string // env file for prod TODO: simplify handling of env files
	EnvFileTest          string // env file for testing
	EnvFileLocal         string // ...
	EnvProd              string // "PROD"
	EnvTest              string // "TEST
	EnvLocal             string // "DEV"
	MaxTweets            int    // maximum number of tweets to get per request
}

func Const() Constants {
	return constants
}

var constants = Constants{
	EndpointUserByID:     baseURL + version + getUserByID,
	EndpointTimelineByID: baseURL + version + timeLineByID,
	UsersFileLocal:       "./users.json",
	UsersFileCloud:       "./serverless_function_source_code/users.json",
	EnvFileProd:          ".env",
	EnvFileTest:          "../.env.example",
	EnvFileLocal:         ".env.dev",
	EnvProd:              "PROD",
	EnvTest:              "TEST",
	EnvLocal:             "DEV",
	MaxTweets:            20,
}
