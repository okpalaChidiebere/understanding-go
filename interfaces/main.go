package main

import "fmt"

/*
We are using a struct here so that we can have a function associated with these types
to demonstrate the importance of interfaces :)
*/
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting1(eb)
	printGreeting(sb)
}

func (eb englishBot) getGreeting() string {
	//VERY custom login got generating an english gretting
	//In a real situation, imagine that the login for getting gerting in english will be very different for the function for getting greeting in spanish
	return "Hi there!"
}

//NOTE: that for the receiver we just wrote the type and not variable. This because we are not using the receiver variable in our function. So it becomes optional to add the variable in the reciever so we just left the type :)
func (spanishBot) getGreeting() string {
	//VERY custom login got generating an spanish gretting
	//...
	return "Hola!"
}

/*Notice that printGreeting1 and printGreeting methods has thesame logic!
With Interfaces you can avoid this and reuse thesame logic for both types englishBot and spanishBot
*/
func printGreeting1(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

/**/
