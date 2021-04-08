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

	/*Although this is a for loop, i will not be interating over and over again millions of times per second. The reciver value "<-c" is still s blocking operation
	The for loop will only loop anytime we receive a value through the channel*/
	for { //we want the loop to run forever so, we deleted the "i := 0; i < len(links)-1; i++ "
		//we refetch the link being fetched again
		go checkLinkAsync(<-c, c) //first argument will be the link that the previous routine returned(go know the type fo value returned by the channel is a string), the second argument will be the channel
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
