// Package utils provides some simple utility methods
package utils

import (
	"doescher.ninja/twitter-service/config"
	"encoding/json"
	"log"
	"os"
)

// DataSource represents a data source.
type DataSource struct{}

// DataSourceInterface is an interface that defines the behavior for retrieving profile IDs from a data source.
// GetProfileIDs returns a list of profile IDs.
type DataSourceInterface interface {
	GetProfileIDs() []string
}

// GetProfileIDs reads profile IDs from the data source.
// It calls the readFromSource function to read the data source content.
// It then unmarshals the content into a slice of strings representing user IDs.
// If there is any error during the process, it calls the FatalIfError function to crash the program.
// Finally, it returns the slice of user IDs.
func (d *DataSource) GetProfileIDs() []string {
	dataSourceContent, err := readFromSource()
	FatalIfError(err)

	var userIDs []string
	err = json.Unmarshal(dataSourceContent, &userIDs)
	FatalIfError(err)

	return userIDs
}

//
//// ReadUserIDs reads a data source and returns a slice returning the individual IDs
//func ReadUserIDs() []string {
//	dataSourceContent, err := readFromSource()
//	FatalIfError(err)
//
//	var userIDs []string
//	err = json.Unmarshal(dataSourceContent, &userIDs)
//	FatalIfError(err)
//
//	return userIDs
//}

func readFromSource() ([]byte, error) {
	return os.ReadFile(config.Conf().FilePath)
}

// FatalIfError crashes the program on error
func FatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
