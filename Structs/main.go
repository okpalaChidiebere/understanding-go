package main

import "fmt"

type ContactInfo struct {
	email   string
	zipCode int
}

/*
Define Exactly what different fields this struct has
Notice that we did not separate the fields by commas, columns, etc Just property names and their types
*/
type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

/*Person and Person2 are still thesame*/
type Person2 struct {
	firstName   string
	lastName    string
	ContactInfo //shorthand way of embedding a struct into another struct. ES6 in JS has this feature. yiu can imagine this as  "ContactInfo ContactInfo". So ContactInfo field of ContactInfo type
}

func main() {

	/*
		Embedding one struct inside of another

		Assume a Person Struct has a Contact Information. This contact information
		contains a Zip code and email address.

		We created a ContactInfo Struct and embed it into the Person Struct

		NOTE: in go, you must habe a comma at the end of each filed of the struct.
		It is is not thesame in JS where you can skip the comma becuase it is the last field of the struct
	*/
	user := Person{
		firstName: "Chidi",
		lastName:  "Okpala",
		contact: ContactInfo{
			email:   "chidi@gmail.com",
			zipCode: 9000,
		},
	}

	user2 := Person2{
		firstName: "Chidi",
		lastName:  "Okpala",
		ContactInfo: ContactInfo{
			email:   "chidi@gmail.com",
			zipCode: 9000,
		},
	}

	fmt.Printf("%+v\n", user)  //Prints : {firstName:Chidi lastName:Okpala contact:{email:chidi@gmail.com zipCode:9000}}
	fmt.Printf("%+v\n", user2) //Prints : {firstName:Chidi lastName:Okpala contact:{email:chidi@gmail.com zipCode:9000}}
}
