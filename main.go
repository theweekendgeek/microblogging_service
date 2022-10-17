package main

import (
	_ "doescher.ninja/twitter-service/orm"
	t "doescher.ninja/twitter-service/twitter"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	t.GetData()
}
