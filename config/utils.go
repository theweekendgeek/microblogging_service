package config

import (
	"encoding/json"
	"log"
	"os"
)

func IsProd() bool {
	return setEnvironment() == Const().EnvProd
}

func IsDev() bool {
	return setEnvironment() == Const().EnvLocal
}

func isTest() bool {
	return setEnvironment() == Const().EnvTest
}

func ReadUserIDs() ([]string, error) {
	readFile, err := os.ReadFile(Conf().FilePath)
	FatalIfError(err)

	var userIDs []string
	err = json.Unmarshal(readFile, &userIDs)
	FatalIfError(err)

	return userIDs, err
}

func FatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
