package main

import (
	"errors"
	"net/http"

)

type requestResult struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request failed")

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://www.airbnb.co.kr/",
		"https://www.google.com/",
		"https://academy.nomadcoders.co/",
		"https://github.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
	}
	for _, url := range urls {
		go hitURL(url, c)

	}
	for i := 0; i < len(urls); i++ {
    requestResult:=<-c
		results[requestResult.url] = requestResult.status
	}
}
func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	} else {
		c <- requestResult{url: url, status: status}

	}

}
