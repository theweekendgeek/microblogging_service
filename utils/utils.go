package utils

import (
	"doescher.ninja/twitter-service/config"
	"encoding/json"
	"log"
	"os"
)

// ReadUserIDs reads a data source and returns a slice returning the individual IDs
func ReadUserIDs() ([]string, error) {
	dataSourceContent, err := readFromSource()
	FatalIfError(err)

	var userIDs []string
	err = json.Unmarshal(dataSourceContent, &userIDs)
	FatalIfError(err)

	return userIDs, err
}

func readFromSource() ([]byte, error) {
	return os.ReadFile(config.Conf().FilePath)
}

// FatalIfError crashes the program on error
func FatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
