/*
To run test in go, you must create a filename that ends in "_test.go".
Basically with go to run test you just have to run write your code in go. It
is not like in JS where you have to install Mocha library

Example of how to test a function
For our newDeck() method, we might want to decide if the length of the deck
returned is equal to a certain number or if the first card in the deck is
equal to a certain number

Basically what to have in mind when you are testing your method is what you
care about in the function as it run.

use "go test" command to run your test file(s). If there ins not test file, nothing will happen when you run this command
*/

package main

import "testing"

/*
The "t" is a testing hanlde that we tell the test when something goes wrong with our testing
Later we will learn what the "*" symbol is
*/
func TestNewDeck(t *testing.T) {
	d := newDeck() //REMEMBER: you can call this method here because this test file belongs to thesame package as deck.go

	//We expect the total number of cards in the deck to be 16. If it not 16, we know something went wrong
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	//Testing that the first card in the new deck is "Ace of Spades"
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	////Testing that the last card in the new deck is "Four of Clubs"
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}

	/*
		When you test this function, go doesn't know that you wrote three tests. This is why you will see
		Just one pass when you test this function. All what go testing does is listen to when one of the error handler "t.Errorf()"
		is invoked so then will it so a fail test message otherwise you will see just one message for
		PASS test not 3 test passed
	*/
}
