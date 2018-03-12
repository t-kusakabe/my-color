package main

import (
	"code.google.com/p/go.net/html"
  "fmt"
  "io"
)

func Scraping(url string) []Result {
  res, err := http.Get(url)
  if err != nil {
    fmt.Println(err)
  }

  defer res.BodyClose()
  results := ParseItem(res.Body)

  return results
}

func main() {
  url := "https://github.com/users/t-kusakabe/contributions"
  dom := Scraping(url)
}
