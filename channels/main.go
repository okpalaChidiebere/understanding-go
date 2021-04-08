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

	/*
		You noticed that the order we made the request, will be the order that gets printed out; In a Serial order
		As well as there was a delay on a few of the links as requests were being made. Reaslistically this is not the best approach.
		You dont have to make a request and wait around for it to finish before you make check for the next link
	*/
	/*for _, link := range links {
		checkLink(link)
	}*/
	/********************************************************************/
	/*
		Better approach is to make these request in parallel instead of serial like the checkLink() method does
		We make a request to each one, and whenever any of them return a response, we then print the response
		This is where Go Routines and channels come to play
	*/

	//We declare our channel and the type of communication we want to communicate over this channel
	//Without the channel, a Parent routine or thread will not know when its child thread or routine finishes. You you have your code "returning early" because it does not expects a communication from a go routine
	c := make(chan string) // communitate through the channel values of type string. Remeber the make function is used to initialize a variable to empty if you are using the shortHand way to initalize variable which is :=

	for _, link := range links {
		//Anytime we use the go keyword we create a background thread(child or go routine) to run the code in out function
		go checkLinkAsync(link, c) //We have to pass the channel into the function so that this function running in the background thread will have to ability to communicate back to the main thread
	}

	//one way to receive message from a channel. We can pass it into a function directly like for this print function
	//another way is to pass it to  varable like myValue := <- c
	//https://stackoverflow.com/questions/37194379/golang-goroutine-cannot-use-function-return-values
	//Imagine Receiving messages like "await" in JavaScript. The Main thread will not go any futher, it will await results from the background thread before it continue
	//fmt.Println(<-c) //NOTE: receiving message from a channel is a blocking line of code like http.Get()

	/*If we dont wait for the other go routine() to emmit messages or results from the get requests in the main thread, we will have our code exit early and the other request will stop running.
	To avoid thwe we have to wait for all channels to complete*/
	/*fmt.Println(<-c) //wait for the second get request result. we are basically telling the compiler, "wait, there is more result we expect to use from the background thread"
	fmt.Println(<-c) //wait for the third get request result
	fmt.Println(<-c) //wait for the fourth get request result
	fmt.Println(<-c) //wait for the fifth get request result*/

	//Having to manually tell our program to await might be too much expecially when the task you want to run asynschronously get too much. Sometimes you might loose count of the thread you created yourself yourself
	//A better way is to have a for loop that will sensure all the background thread completes before the program exits. You can imagine this as Promise.all() method for JavaScripts ES6
	for i := 0; i < len(links)-1; i++ {
		fmt.Println(<-c) //each time this print line is called, it block the main thread from execting, so we can still get all our results back
	}

	/*
		I my commadLine results looks like

			http://golang.org is up!
			http://golang.org
			http://google.com is up!
			http://google.com
			http://stackoverflow.com is up!
			http://stackoverflow.com
			http://amazon.com is up!
			http://amazon.com
			http://facebook.com is up!
			http://facebook.com

			Notice that the print results are not in serial order anymore like the way the are arranged in the slice
	*/

}

//takes a link and makes a get request to it and respond whether the URL is responsing to traffic or not
func checkLink(link string) {
	_, err := http.Get(link) //we only care about the error variable. We dont need the respnse from the request. Thats why we have an underscore for it. This will make the compiler not complain as an unUsed variable

	//if there an actual value in the error variable, then the website might be down
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
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
