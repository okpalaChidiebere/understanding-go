package main

import "fmt"

//you can image this as defining your class like in c++, javascript or java
type deck []string //create a new type of 'deck' which is a slice of strings

/*
Creates and returns a list of playing cards
Imagin this method as your "constructor" for C++, Java or Javascript in OOP

We will not add any receiver to this function because, chances are when you are
calling this methid, you dont have an exsiting deck. You want a new deck defined
*/
func newDeck() deck {
	cards := deck{} //creates a new variable of type deck which is empty

	//To avoid manually having to type out all 15 different combinations of cards we will use two for loops to
	//to initialize our empty deck
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"} //defining list of card suits
	cardValues := []string{"Ace", "Two", "Three", "Four"}          //defining a list of card values
	for _, suit := range cardSuits {                               //NOTE: we use underScore(_) instead of the index to get around the syntax error of not using a defined index variable
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards //returning our new deck
}

//You can imagine this as defining the methods that belongs to your class like in c++, javascript or java
func (d deck) printCards() { //NOTE: the "(d deck)" is a receiver.

	//you can see the 'd' (the receiver argument) similar to thesame keyword "this" for Java or Javascript. 'd' is a value of deck in our case
	//In go, there is no 'this' or 'self' keyword, we will always use the actual value that it is instead of "this" or "self"
	//for naming converntion, a good practise is to have the 'd' which is the first letter of name of the custom type or a two or three letters anming convention like 'dec'; but you can call it whatever you want!
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/*
Go is not an object oriented programming Language like C++, Java, Python. So, there is no idea of classes inside of Go

Eg in Real OOP, we could have a Deck class of which we can create an instance of that Deck class where we can have some
functions defined inside the scope like print, shuffle, etc that will be assigned to any instance of the Deck class

But if you were to approach this in Go world, it is quite different.
Instead of defining a class, you will define a new type called Deck
Now, to define the custom functionality to work with your custom type, you attach a receiver to it. A function with a
receiver is in go is thesame to a function that belongs to a class or instance in OOP.

*/
