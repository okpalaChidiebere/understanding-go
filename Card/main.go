package main

func main() {

	cards := newDeck() //REMEMBER: we can call this method in the deck.go file from this file, because main.go and deck.go belongs to thesame package "main"!

	/*
	  we can now call the printCard() method on the cards variable. This is because of the receiver we set up on it

	  Receiver adds methods on the variables that we create of that type
	*/
	cards.printCards()
}
