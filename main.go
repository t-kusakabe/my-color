package main

import (
	. "./github"
	. "./twitter"
	"fmt"
)

func main() {
	grassHex := ScrapingGrass()

	api := Credentials()
	tweet, err := api.PostTweet("My color today is" + grassHex, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tweet.Text)
}
