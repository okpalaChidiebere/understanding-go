package main

//When you have more than one import, you wrap your libraries in a parenthesis with no commas between them
import (
	"fmt"
	"io/ioutil"
	"os"      //this package works well across all operating systems your code is running be it windows, mac, linux
	"strings" //more on this library here here https://golang.org/pkg/strings/
)

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
The idea behind the deal method is that we take an existing deck of cards,
we will take a number of cards to deal out and we will create a new deck of
the specified size. This will simulate having a deck of 52 cards in practise
and dealing out a hand of 5 to 7 cards. Whatever size you want.

Eg we have deck that has 4 cards, and we specify that we want  deals(3)(3 cards in our hands)
and 1 of our card on the deck, it will return two two decks of one that contains
3 cards and a deck that contain one card

If we want to select a range within a specific slice
fruits[startingIndexIncluding : upToAndNotIncluding]
Eg assume we have fruits["apple", "banana", "orange", "mango"]
fruits[0:2] or fruits[:2] gives us fruits["apple", "banana"]
fruits[2:] gives us fruits["orange", "mango"]

First argument: is our deck. Remember naming convention for go is to have the 1 to 3 characters for the actual variable.
There is no this or self keyword in go

Second Argument: is amount of cards we want in our hands that we want to return

Return value: The function returns two values of type deck. In go you can tell the compiler that we are returning two values. C++ has this functionality as well
*/
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:] //we return the first deck which is the one in our hand, and the second value is the remaining cards on deck
}

/*It takes a deck and returns a complete representation of that deck in string

NOTE: we attached a deck receiver to this function. Now, you can optionally pass the deck as an argument
But for our project we wanted to be able to call cards.toString() so having a receiver makes more sense.
Down the line you will learn when to use a receiver over a argument for you custom type
*/
func (d deck) toString() string {

	/*
		TypeConversion in Go is where you take one type value and convert it to another type. To achieve
		this in go you; have the type you want on the left side, a parenthesis and the type that you have
		inside the parenthesis. Eg: []string(d) we have our deck and we convert it to a slice or array of strings
	*/
	return strings.Join([]string(d), ",") //We joined every string at each index in the array to one string with a comma between each value.
}

/*Again for this function signature we want this function to have a receiver so that we can call
cards.saveToFile(). As well as we will be converting the deck to string d.toString() inside this function as well

We return the error because it is better rather than us trying to figure out what is going on with some type of error being generated by writing something to a file
we return the error that might bet produced when we attempt to write something to a hardDrive
*/
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) //0666 is apermission where anyone can read or write to this file
	//this writeFUnction could return an error according to documentation so we return the whole result of this write function
}

/*
We dont need any receiver for this function because we dont yet have a deck
when this function is called
*/
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) //the ReadFile method returns two values slice of bytes which is a slice of strings and error value

	//If something when wrong during the read process the err variable will be populated with an error value return by the readFile method otherwise it will be nil
	if err != nil { //error handling
		//The two options below is what you as as developer can implement in a real situation when something goes wrong while we get something from a file or some logi in your code
		// Option #1 - log the error and return a call to newDeck(). At least we still give the user a deck
		// Option #2 - Log the error and entirely quit the program
		fmt.Println("Error:", err) //log the error
		os.Exit(1)                 //quit the program
	}

	/*
		Recap ****
		We have a deck and we want to eventually turn it into a byte of slices
		So we have our deck which is a slice of string, turn it to a slice of strings,
		then join all the slices into a single string, then take that string and turn
		it into a byte slice then write the bytes to hard drive

		To get back our deck from the textFile in our drive, we basically did the
		reverse process of what we did to save it. We read the bytes from the file
		in the drive, perform a type cast that converts the bytes to string, then
		perform a split operation in the string split the string into an array,
		then finally back to our custom type deck
	*/
	s := strings.Split(string(bs), ",") //NOTE: you will get use to seeing very short variable names in go like the "s" :)
	return deck(s)
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
