package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	urls := getUrls()

	c := make(chan string)

	for _, url := range urls {
		go checkUrl(url, c)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<- c)
	}

}

func getUrls() []string {
	urls := []string{
		"https://young-springs-45765.herokuapp.com/hello",
		"https://young-springs-45765.herokuapp.com/materials",
	}

	for i := 1; i < 10; i++ {
		url := "https://young-springs-45765.herokuapp.com/post?postId=" + strconv.Itoa(i)
		urls = append(urls, url)
	}

	return urls
}

func checkUrl(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- "Site is down: " + link

		return
	}

	c <- "Site is up: " + link
}
