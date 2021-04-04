package main

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
	cards := deck{newCard(), "Five of Diamonds"} //card is of card deck
	cards = append(cards, "Six of Spades")       //adding a new element to the Slice. The append method returns the new slice that is assgined to the cards variable

	/*
	  we can now call the printCard() method on the cards variable. This is because of the receiver we set up on it

	  Receiver adds methods on the variables that we create of that type
	*/
	cards.printCards()
}

/*We will define a new function that will return a string of "Ace of Spaced"

you can do this in main function card := newCard() */
func newCard() string { //note: you MUST specify the return type, if your function will return a value
	return "Ace of Spaces"
}
