package main

import (
	"doescher.ninja/twitter-service/orm"
	t "doescher.ninja/twitter-service/twitter"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = orm.Connect()
	if err != nil {
		log.Fatal(err)
	}

	err = orm.Migrate()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	t.GetProfiles()
}
