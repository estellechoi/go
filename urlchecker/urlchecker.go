package urlchecker

import (
	"fmt"
	"net/http"
)

type hitResult struct {
	url    string
	status int
}

// chan<- means sendonly, cannot receive
func hitURL(url string, c chan<- hitResult) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		c <- hitResult{url: url, status: 0}
	}
	c <- hitResult{url: url, status: 1}
}

func main() {
	// url checkers
	c := make(chan hitResult)

	results := map[string]int{}
	urls := []string{"https://www.airbnb.com", "https://www.google.com", "https://soundcloud.com"}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i += 1 {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}

}
