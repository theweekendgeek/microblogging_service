package service

import (
	"doescher.ninja/twitter-service/business"
	// load config
	_ "doescher.ninja/twitter-service/config"
	"doescher.ninja/twitter-service/persitence"
	"fmt"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"net/http"
)

func init() {
	functions.HTTP("HelloGet", HelloGet)

	persitence.InitDatabase()
}

func HelloGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("INVOKING FUNCTION")
	business.GetNewTweets()
	fmt.Println("FUNCTION DONE")
}
