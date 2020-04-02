package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
		"https://amazon.com",
		"https://github.com",
		"https://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(l string, c chan string) {
	_, err := http.Get(l)

	if err != nil {
		fmt.Println(l, " is down!")
		c <- l
		return
	}

	fmt.Println(l, " is up!")
	c <- l
}
