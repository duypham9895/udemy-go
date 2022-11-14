package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.google.com/",
		"https://www.edwardpham.me/",
		"https://www.linkedin.com/in/phamanhduy",
		"https://github.com/duypham9895",
	}

	c := make(chan string)

	for _, url := range links {
		go validateURL(url, c)
	}

	for url := range c {
		go func(url string) {
			time.Sleep(time.Second)
			validateURL(url, c)
		}(url)
	}
}

func validateURL(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url + " might be down!")
		c <- url
		return
	}
	fmt.Println(url + " is up!")
	c <- url
}
