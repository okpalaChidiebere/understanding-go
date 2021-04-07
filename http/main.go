package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	/*
		Having this is not enough. We could not find the body of the response.
		In other programming languages like JS this would be fine. But with go
		its different. We need the help of interfaces

		Looking at the documentation https://golang.org/pkg/net/http/#Response, we were logging out a struct

		The body is of type ReadCloser. If it were js, the body would be a JSON or a blob of string
		We now dig into https://golang.org/pkg/io/#ReadCloser to see its an interface!
		The ReadCloser implements the Reader interface
	*/
	fmt.Println(resp)

	/*
		Here we are mainig use of the Read function implemented by the Reader interface to get the bytes raw data that gets converted to bytes

		if you look at this https://golang.org/pkg/io/#ReadCloser you will notice that the Reader interface(https://golang.org/pkg/io/#Reader) is inside ReadCloser interfce.
		Yes, it is because in go, we can wrap an interface inside an interface!

		Because the ReaderCLoser implemets the Reader interface, it is have access to call the Read() method
	*/
	bs := make([]byte, 99999) //We make a asummtion 99999 spaces or elements is big enough to accomodate the data that is returned by our Reader. If we initialize a byte slice with 0 elements inside of it, the Reader will assume the reader is full and will just quit early
	resp.Body.Read(bs)        //We get the raw html from the body field and read that into our byte slice
	fmt.Println(string(bs))   //We converted our bytes to string because the commandline cannot write bytes

	/*
		The process of which we declare a space for bytes, read bytes convert it to string to print to the terminal we dont always have to do
		Go had built in helper functions to do this for us and it is the Copy() function https://golang.org/pkg/io/#Copy

		hold command button and click on the Copy() function to see for yourself its implementation in your IDE

		Ths Copy function expects the
		- first argument to be a value that implements the Writer interface. Rememebr the writer interface takes data and sends it outside of our application
		- The second argument to be a value that implements the Reader inteface
		In summary the Copy function helps takes data from inside of our application and write or copy the information to outside of our app. Basically like the three lines of code above

		//os.StdOut variable is of type File that implements the Write interface because the File type has a function called Write() that receive byte slice and returns an int or error https://golang.org/pkg/io/#Writer
		//resp.Body varibale implements the Read interface already


	*/
	io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}

	/*
		As long as io.Copy() is concern, lw is a value that implements the writer interface because
		satisfies it has a function called Write() that takes a byte slice and returns an int and an error

		If we put together a Junk or crap logic inside the function that does not do the right thing, it will compile and run
		but our program will not do what it is actually meant to do
	*/
	io.Copy(lw, resp.Body)
}

/*
The signature of the function associated to our custom type must satisfy "Write(bs []byte) (int, error)" for it to implement Writer interface
https://golang.org/pkg/io/#Writer

Dont forget why we can just write the type in the receiver without the value is because we are not using the value in our function! :) Its a shorthand way
*/
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
