package main

import "fmt"

//card := "Ace of Spades" this will give a runtime error.Variables can be initialized outside of a function, but cannot be assigned a value.
//Eg: var deckSize int  is valid. We can now later assign value to it inside the main function.
func main() {
	/*
		using the keyword "var" we are telling go compiler we are creating a variable
		We declare varibales in go must assign it a static type just thesame way we do in Java, C++ or Typescript.
		In JS or python we dynamically assign variables

		Basic static types in go are bool, int, float, string. There are many more as well
	*/
	//var card string = "Ace of Spades" //long format way of declaring a varibale
	card := "Ace of Spades"   //shortHand way of declaring a varibel in go. We only use ":=" when we are initializing a new variable
	card = "Five of Diamonds" //we change the value of card. Notice that we did not use the ":" symbole anymore for re-assigning

	/*
			We might be able to initialize a variable and then later assign a variable to it

			Eg:
			var deckSize int
		    deckSize = 52
	*/

	fmt.Println(card)
}

/*We will define a new function that will return a string of "Ace of Spaced"

you can do this in main function card := newCard() */
func newCard() string { //note: you MUST specify the return type, if your function will return a value
	return "Ace of Spaces"
}

/* How you declare a function that return void or nothing in go

func newCard2() {
	return
}

*****************************
If we had another file called state.go with the following code

package main

import "fmt"

func main() {
    fmt.Println(newCard()) //notice that we calling the method newCard from another file called main.go
}

This is valid because Files in the same package can freely call functions defined in other files in go
*/
