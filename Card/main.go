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
