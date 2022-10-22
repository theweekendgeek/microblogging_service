package twitter

import (
	"doescher.ninja/twitter-service/utils"
)

func Request[T any](url string) *T {
	resByte := MakeRequest(url)

	var resObj T
	err := Parser{}.ParseResponse(resByte, &resObj)
	utils.FatalIfError(err)

	return &resObj
}
