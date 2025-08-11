package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://t-t.com",
		"https://lantigua.rulyotano.com",
		"https://google.com",
		"https://rulyotano.com",
		"https://github.com",
		"https://facebook.com",
		"https://nginx.com",
	}

	CheckSites(links)
}

func CheckSites(links []string) {
	c := make(chan string)

	for _, link := range links {
		go CheckLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			CheckLink(link, c)
		}(l)
	}
}

func CheckLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", link, err)
		c <- link
		return
	}
	fmt.Printf("Successfully fetched %s\n", link)
	c <- link
}