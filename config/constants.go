package config

var baseUrl = "https://api.twitter.com"
var version = "/2"

var getUserByID = "/users/"
var timeLineByID = "/users/%s/tweets?max_results=20" //+ fmt.Sprintf("%d", Conf().MaxTweets)

type Constants struct {
	UserByID       string
	TimelineByID   string
	UsersFileLocal string
	UsersFileCloud string
	EnvFileProd    string
	EnvFileTest    string
	EnvFileLocal   string
	EnvProd        string
	EnvTest        string
	EnvLocal       string
}

func Const() Constants {
	return constants
}

var constants = Constants{
	UserByID:       baseUrl + version + getUserByID,
	TimelineByID:   baseUrl + version + timeLineByID,
	UsersFileLocal: "./users.json",
	UsersFileCloud: "./serverless_function_source_code/users.json",
	EnvFileProd:    ".env",
	EnvFileTest:    "../.env.example",
	EnvFileLocal:   ".env.dev",
	EnvProd:        "PROD",
	EnvTest:        "TEST",
	EnvLocal:       "DEV",
}
