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

	user.print() //Prints : {firstName:Chidi lastName:Okpala contact:{email:chidi@gmail.com zipCode:9000}}
	user.updateName("ifeanyi")
	user.print() //**Still Prints : {firstName:Chidi lastName:Okpala contact:{email:chidi@gmail.com zipCode:9000}}. The ifeanyi is not updated in the variable of the main function but inside the updateName() itself it works
	/*To fix having the user variable update in the main function, we use a point to actually pass the reference of the real value in memory and not a copy.*/
	//userPointer := &user //give the memeory address of the value that this variable(user) is pointing to. So now the userPointer variable now points to thesame adderess where the user variable points to. &<variable> turns your value into an address
	//userPointer.updateNameByReference("ifeanyi")
	user.updateNameByReference("ifeanyi") //For shortcut you dont need to wite the & all the time to convert value to maac address and pass the address referecne. As long as you have the "*" symbol next to the type in the receiver or function argument, go will update it properly for you!
	user.print()                          //Prints; {firstName:ifeanyi lastName:Okpala contact:{email:chidi@gmail.com zipCode:9000}} Now we have the varibale updated in our main function

	name := "bill"

	namePointer := &name

	fmt.Println(&namePointer) //prints address: 0xc00000e028
	printPointer(namePointer)
}

func (p Person) print() {
	fmt.Printf("%+v\n", p)
}

/* This method when called will update the person's firstname

By default, this is called by value; meaning a copy of your variable is passed to your receiver
*/
func (p Person) updateName(newfirstName string) {
	p.firstName = newfirstName
	fmt.Printf("%+v\n", p) //prints: {firstName:ifeanyi lastName:Okpala contact:{email:chidi@gmail.com zipCode:9000}}

}

/*NOTE: Passing by reference cannot be only done in the receiver with the pointers, you can have pointer in the arguments for your function as well
But this works for this scenario*/
func (pointerToPerson *Person) updateNameByReference(newfirstName string) {
	/*
		now we have access to the real value of the variable because "&user" gave
		use the MAC address where the variable is stored. If we had not passed the
		&<variableName>, by default a copy of the variable will be made and we will
		not be able to update the variale in the main function. It will work just fine but not just the outcome we expected
		*<address> like *Person turns your address passed into a value. * tells us that we are pointing to a MAC address where a varible is stored
	*/
	(*pointerToPerson).firstName = newfirstName //(*pointerToPerson) because of the brackets around the variable, we now have access real struct value as well as its fields. You can see it as desturcting your value. So now we can update it or do whatever
}

func printPointer(namePointer *string) {
	fmt.Println(*&namePointer) //prints address 0xc00000e038
	//The address is of a different value because by default, *everything* in go is pass by value. But the value can still be referenced because its a pointer
	fmt.Println(*namePointer) //prints: bill as you see we still printed bill
}
