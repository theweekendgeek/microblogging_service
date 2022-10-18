package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

func ReadUserIds() ([]string, error) {
	readFile, err := ioutil.ReadFile(Conf().FilePath)
	FatalIfError(err)

	var userIds []string
	err = json.Unmarshal(readFile, &userIds)
	FatalIfError(err)

	return userIds, err
}

func FatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
