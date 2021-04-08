package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLinkAsync(link, c)
	}

	for l := range c { //so now this for loop will wait for next next channel value. The range c is now a Block code
		//we refetch the link being fetched again
		go checkLinkAsync(l, c) //first argument will be the link that the previous routine returned(go know the type fo value returned by the channel is a string), the second argument will be the channel
	}

}

func checkLinkAsync(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link //send a message into our channel
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
