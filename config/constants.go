package config

var baseUrl = "https://api.twitter.com"
var version = "/2"

var getUserById = "/users/"
var timeLineById = "/users/%s/tweets?max_results=20" //+ fmt.Sprintf("%d", Conf().MaxTweets)

type Constants struct {
	UserById       string
	TimelineById   string
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
	UserById:       baseUrl + version + getUserById,
	TimelineById:   baseUrl + version + timeLineById,
	UsersFileLocal: "./users.json",
	UsersFileCloud: "./serverless_function_source_code/users.json",
	EnvFileProd:    ".env",
	EnvFileTest:    "../.env.example",
	EnvFileLocal:   ".env.dev",
	EnvProd:        "PROD",
	EnvTest:        "TEST",
	EnvLocal:       "DEV",
}
