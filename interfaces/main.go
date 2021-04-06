package main

import "fmt"

type bot interface {
	/*
		We are basically saying, if there is a type in this program or workspace that implements the function
		getGretting() that returns a string is now an honorary member of the type bot. This means those types can now also be of type bot
		Go takes care of the mactching for us. You dont have to write an extra code to say types(englishBot, spanishBot) can be used as bot type

		Being an honourary memeber means that those types can now call the printGreetingWithInterface() method
		This means we can now takes the custom types(englishBot, spanishBot) and pass them down to type bot as well

		Being an honorary mameber is one of the rules of interface

		With inteface you CANNOT create a value directly  out of interface type like you do for concrete type like int, string, map, struct or your own custom type that extends the standard type like englishBot etc

		You will start to see terms like concrete type or interface type in documentations in go so remember what they are!
	*/
	getGreeting() string //Note we are not only limited to define the return type of a function but also the argument types.
	/*
		You can list out multiple different functions that a type has to satisfy to qualified as being of type bot. Eg
		getGreeting(int, string) (string, error)
		getBotVersion() float64
		respondToUser(user) string
	*/
}

/*
We are using a struct here so that we can have a function associated with these types
to demonstrate the importance of interfaces :)
*/
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreetingWithInterface(eb) //prints: "Hi there!"
	printGreetingWithInterface(sb) //prints: "Hola!"
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

func printGreetingWithInterface(b bot) {
	fmt.Println(b.getGreeting())
}
