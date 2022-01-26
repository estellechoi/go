package main // to compile

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/estellechoi/go/accounts"
	"github.com/estellechoi/go/dict"
)

func main() {
	// accounts
	account := accounts.NewAccount("Yujin")
	fmt.Println(account)

	account.Deposit(1000000000000000)
	fmt.Println(account)
	fmt.Println(account.Balance())

	err := account.Withdraw(1000000000000000000) // error handling in Go (strong error checking)
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
	fmt.Println(account.Owner())
	fmt.Println(account)

	fmt.Println("========================================")

	// dict
	dictionary := dict.Dictionary{"name": "Yujin"}
	err2 := dictionary.Add("city", "Donghae")
	if err2 != nil {
		fmt.Println(err2)
	}

	value, err3 := dictionary.Search("city")
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(value)
	}

	err4 := dictionary.Add("city", "Donghae")
	if err4 != nil {
		fmt.Println(err4)
	}

	err5 := dictionary.Update("city", "Seoul")
	if err5 != nil {
		fmt.Println(err5)
	}

	updatedValue, _ := dictionary.Search("city")
	fmt.Println(updatedValue)

	dictionary.Delete("city")

	searchedValue, err6 := dictionary.Search("city")
	if err6 != nil {
		fmt.Println(err6)
	} else {
		fmt.Println(searchedValue)
	}

	fmt.Println("========================================")

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

var errRequestFailed = errors.New("request failed")

func hitURL(url string) error {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}

	return nil
}
