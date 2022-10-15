package main

import (
	t "doescher.ninja/twitter-service/twitter"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	t.Request()

}
