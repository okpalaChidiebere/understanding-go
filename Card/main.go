package main

import "fmt"

/*
Go has two ways of handling list of records

Arrays: Must have Fixed length. Thesame way with C++. As you initialize the array, you define the constant fixed length
Slices: This is an array that can grow or shrink. This one the length may be unknown meaning that it can grow or shrink

Slices and Array must be defined with a dataType
*/
func main() {
	/*We are declaring a Slice of string and assigning records to it right away
	If we dont have any list yet, we could just have it as "{}" meaning empty list of strings
	*/
	cards := []string{newCard(), "Five of Diamonds"}
	cards = append(cards, "Six of Spades") //adding a new element to the Slice. The append method returns the new slice that is assgined to the cards variable

	fmt.Println(cards) //this will print [Ace of Spaces Five of Diamonds Six of Spades]

	/*How to loop through a Slice to access individual record of the slice
	- "Range" is a keyword used when we want to interage over every single record inside of a slice
	- The reason we are using the ":=" syntax over "=" is because with for loop in go, we are throwing away the
	previous index so we need to re-declare for the new index value to card
	- Slice can have only one type of value in it*/
	for i, card := range cards {
		fmt.Println(card, i) //we print the value and the index position of the value in the array
		//SIDE NOTE:for this for loop, we must use the index otherwise you will get a syntax error of unsied variable
		//We'll see a way to deal with the unused 'index' variable later
	}
}

/*We will define a new function that will return a string of "Ace of Spaced"

you can do this in main function card := newCard() */
func newCard() string { //note: you MUST specify the return type, if your function will return a value
	return "Ace of Spaces"
}
