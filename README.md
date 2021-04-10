# LANGUAGE REVIEW

# Variables
You delcare variables using the `var` keyword or the scope variable clarator `:=`
```go
package main

var q int //Package level declaration ONLY, put you can use it for block level claration as well. The "var" keywords sets the value of the varibale to zero value 
func main(){
   //We have short variable declaration operator which you should try to use as much as possible. You only use this in a code block(code wrap around { })
   x := 7 //block level scope declaration of variables only
   fmt.Println(x) //prints 7
   fmt.Printf("%T", x) //prints the type of the operator using %T. wpe get int printed out
   q = 8 

}
```

# Data Structures

NOTE: for structs if you want to initialize the values right away you do it in a composite literal way. This means where after the ":=" you have the type then the {} with values inside the parenthesis

`slice` -  is a data structure for holding a list of things or types
`maps` - are for key values and quick look ups by key
`struct` - This is equivalent of an object in OOP languages where we have fileds and they store something

There is also "for range" loop for mapping through list. Check that out in this git 

```go
package main

type Person struct {
    fname string //If you want the filed to be visible outside the package, then the field fist letter must start with uppercase eg Fname
    lname string
}

func main(){
    //declaring a list of integers
    xi := []int{2,3,4,5,42} //when you list things in one row or horizontally, you dont need the trailing comma at the end. You can have it if you want it
    fmt.Println(xi)

    //declaring a map
    m := map[string]int{
        "scot": 41, 
        "jon": 42, //when you list things vertically, you do need the trailing comma at the end
    }
    fmt.Println(m)

    //declaring a struct
    p1 := Person{ //since we are putting the field values in order, we dont need to provide the field name
        "Miss",
        "MoneyPenny",
    }

}
```

# Functions
The signature looks like
func (receiver) identifier(parameters)(return) { <code> }

Note the reciver is optional but is used to attach a function of an identifier to a type

```go

package main

type Person struct {
    fname string
    lname string
}

//SecretAgent extends the person type and have extra attributes as well
type SecretAgent struct{
    Person
    licesnseToKill bool
}

//This function returns nothing
func (p Person) speak{
    fmt.Println(p.fname, p.lname,`says, "Good Morning, James."`)
}

//A secret agent method
func (sa SecretAgent) speak{
    fmt.Println(sa.fname, sa.lname,`says, "Shaken, not Stirred."`)
}

func main(){
    p1 := Person{ 
        "Miss",
        "MoneyPenny",
    }

    p1.speak() //prints: Miss MoneyPenny says, "Good Morning, James."

    //SecretAgent is a Person and more..
    sa1 := SecretAgent{
        Person{
            "James",
            "Bond",
        },
        true
    }
    sa1.speak() //prints: James Bond says, "Shaken, not Stirred."
    //You can drill down to the inner type and call fields or methods associated with that inner type. In our case the inter type is "Person"
    sa1.Person.speak() //prints: James Bond says, "Good Morning, James."
}
```

# Interfaces and Polymorphism

An interface is a defined functionality and allows polymorphism. Basically if we have two values of different types but they implement the interface, then they are also of the interface type

```go

package main

type Person struct {
    fname string
    lname string
}

type SecretAgent struct{
    Person
    licesnseToKill bool
}

//Person and SecretAgent can also be a Human. 
type Human interface {
    speak() //Humans can speak
}


func (p Person) speak{
    fmt.Println(p.fname, p.lname,`says, "Good Morning, James."`)
}
func (sa SecretAgent) speak{
    fmt.Println(sa.fname, sa.lname,`says, "Shaken, not Stirred."`)
}
/*Since both Person and SecretAgent has the speak method. This makes these two types to implements without any explicit declaration the Human interface*/


func saySomething(h Human){
    h.speak() //h for Huma will call their speak method
}

func main(){
    p1 := Person{ 
        "Miss",
        "MoneyPenny",
    }

    saySomething(p1) //prints: Miss MoneyPenny says, "Good Morning, James."

    //SecretAgent is a Person and more..
    sa1 := SecretAgent{
        Person{
            "James",
            "Bond",
        },
        true
    }
    saySomething(sa1) //prints: James Bond says, "Shaken, not Stirred."
    saySomething(sa1.Person) //prints: James Bond says, "Good Morning, James."
}
```