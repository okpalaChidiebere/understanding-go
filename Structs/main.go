package main

import "fmt"

/*
Define Exactly what different fields this struct has
Notice that we did not separate the fields by commas, columns, etc Just property names and their types
*/
type Person struct {
	firstName string
	lastName  string
}

func main() {
	/*One of the several different ways to declare a new struct in go

	GO will assume that the firstname is "Chidiebere".
	So in this scenario, we are relying 100% on the order of our fields for
	assignment of values

	This might not be the best way to declare your struct, but some official code documentation in go uses it
	If we accidentally swap firstName and Lastname in the Person struct definition now we get a different result of
	firstname being Okpala
	*/
	user := Person{"Chidiebere", "Okpala"} //Create a new value of type for the struct you defined

	/*This is a better syntax because we can always change the order of the properties
	in the struct definition and the values will still be thesame.

	This is not possible with the way we declared struct for the "user" variable*/
	anotherUser := Person{firstName: "Ifeanyi", lastName: "Okpala"}

	/*Third way of declaring a new struct.

	With this, since we did not initialize a value to this varibale yet, go will assign a 0 value to the fields

	Here is a list of Zero values for common types
	int: 0, float: 0, bool: false, string: "". These will be default values of type for your variables if you dont assign them any value yet
	In Javascript the default values for un initialized variable is undefined, with go the defualt values is what i listed above.
	However, with struct you may want to define your own default values. It s up to you. If ou dont defined your own, go will assume the default value based on the type for you
	*/
	var thirdUser Person

	//Just like stings we can easily print structs out in the commadline
	fmt.Println(user, anotherUser, thirdUser) //prints in the commandLine: {Chidiebere Okpala} {Ifeanyi Okpala} { }

	//Notice for the third person, the value is being defined but they are both empty stings
	//The use PrintF to confirm that the struct fields are assigned but are empty
	fmt.Printf("%+v\n", thirdUser) //prints to commandline: {firstName: lastName:}

	/*Updating an existing declared struct*/
	thirdUser.firstName = "Chiamaka"
	thirdUser.lastName = "Okpala"
	fmt.Printf("%+v\n", thirdUser) //prints to commandline: {firstName:Chiamaka lastName:Okpala}

}
