package main

import (
	"fmt"
	"net/http"
	"time"
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
		//we run our function literal in a go routine (background theread)
		go func(link string) { //We get all the arguments passed into a function literal here. In our case, only the link returned through the channel was passed
			time.Sleep(5 * time.Second)
			checkLinkAsync(link, c) //first argument will be the link that the previous routine returned(go know the type fo value returned by the channel is a string), the second argument will be the channel
		}(l) /* You have to place an extra set of parenthesis to actually invoke it

		Notice that we passed the link message returned through the channel into the parenthesis. This is a way to pass argument into your function literal in go(Remember this is bass by value which is default in go). You can relate to this as the way React hooks works.
		This is necessay because anytime we are creating a child routine we want to pass the value from the main routine into the function that will run as a child routine

		If we did not pass the l into the function and just call it directly from the checkLinkAsync method, Go will give us a warning that we are referecning to a global variable. We never want to point to thesame variable from different child routine. This is a NO NO
		*/
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
