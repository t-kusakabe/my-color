package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
  . "./twittercredentials"
)

func main() {
	doc, err := goquery.NewDocument("https://github.com/users/t-kusakabe/contributions")

	if err != nil {
		fmt.Println("Can not URL.")
	}

	result, _ := doc.Find("svg g g").Last().Find("rect").Last().Attr("fill")
	fmt.Println(result)

	api := TwitterCredentials()
	tweet, err := api.PostTweet("Test", nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tweet.Text)
}
