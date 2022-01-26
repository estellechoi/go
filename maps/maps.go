package maps

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("request failed")

func hitURL(url string) error {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}

	return nil
}

func main() {
	// url checkers
	results := make(map[string]string) // if you dont want to initialize map, use make()
	urls := []string{"https://www.airbnb.com", "https://www.google.com", "https://soundcloud.com"}

	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAIL"
		}
		results[url] = result
		// fmt.Println(url, err)
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}
