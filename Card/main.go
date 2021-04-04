package main

func main() {

	cards := newDeck() //REMEMBER: we can call this method in the deck.go file from this file, because main.go and deck.go belongs to thesame package "main"!

	//hand, remainingCards := deal(cards, 5) //we defined two separate variables to receive the two separate values that is returned by this function
	//the first return value will be assinged to 'hand', and the second return value will be assigned to 'remainingCards'. Both values of type deck\

	//REMEMBER: you can call the print function on this two variables because they are of type deck
	//hand.printCards()
	//remainingCards.printCards()

	cards.shuffle()

	/*
	  we can now call the printCard() method on the cards variable. This is because of the receiver we set up on it

	  Receiver adds methods on the variables that we create of that type
	*/
	cards.printCards() // After calling "deal" and passing in "cards", the list of strings that the "cards" variable point did not change. We never directly modified the slice that 'cards' is pointing at.
}
