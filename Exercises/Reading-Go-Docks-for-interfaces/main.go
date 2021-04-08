package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {

	//os.Args is a slice ot type string that represents all the commandLine argument passed to our program
	fmt.Println(os.Args) //in your termianal if you run go run main.go blobA blobB our code prints: [/var/folders/v_/dfxq4fkx52gg90wn4n7b7c_r0000gn/T/go-build2582685340/b001/exe/main blobA blobB]. The first element we dont care about
	//The frist element in the slice is the actual executable file that represents our main.go process. Remember that you can create this executable file yourself with "go build" but side we used "go run", go create a temporary one just to run our code and delete it later

	fn := os.Args[1] //gets the fileName passed from the command line that we need to open
	var buf bytes.Buffer

	f, err := os.Open(fn) //open the file in our directory or hardDrive matching the name and read the contents

	//error checking
	if err != nil {
		fmt.Println("Error:", err) //log the error
		os.Exit(1)                 //there is nothing else to do if there is an error reading the file. So we just return early and exit the program
	}

	/*
		since the f is of type File, File(https://golang.org/pkg/os/#File) has a function called Read(keep scrolling down the doc for type FIle and you will see it)
		func (f *File) Read(b []byte) (n int, err error) which is thesame to what is type is qualified as a Reader interface https://golang.org/pkg/io/#Reader

		So we can pass in the f into the second argument of type Reader for the copy function
	*/
	io.Copy(&buf, f)
	fmt.Println("previous output:", buf.String())

	/*
		More easier way to print
		This is recommended the the first way because, the Stdout will guarantee to print on whatever operationg system we run this go program

		Or you can create your own logger like we did for the http class exercise
	*/
	io.Copy(os.Stdout, f)
}

/*
One cool thing is that you can also print your go file
First build your program using "go build main.go"
Then run "./main main.go"
*/
