package main

import (
	"doescher.ninja/twitter-service/business"
	_ "doescher.ninja/twitter-service/config"
	"doescher.ninja/twitter-service/persitence"
)

func init() {
	persitence.InitDatabase()
}

func main() {
	business.RequestAndSaveTweets()
}
