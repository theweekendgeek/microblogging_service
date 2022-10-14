package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}

	username := "BarackObama"
	req, err := http.NewRequest("GET", "https://api.twitter.com/2/users/by/username/"+username, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+os.Getenv("BEARER"))
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Body)
}
