package main

import "fmt"

func main() {

	/*
		We are declaring a map variable where thet all of the keys will be of type string and the values will be of type string as well

		We want this map, to get the hexCode for a color given the key
	*/
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	rainbowColors := make(map[string]string) //This works also. But we usually use this when we want to later decide what keys and values to add to this map

	fmt.Println(colors)        //prints:  map[green:#4bf745 red:#ff0000 white:#ffffff]
	fmt.Println(rainbowColors) //prints: map[]

	rainbowColors["green"] = "#4bf745"  //adds a new keyValue pair to your empty map
	fmt.Println(rainbowColors["green"]) //prints: #4bf745 Notice that with maps we dont use the dot(.) syntax like struct in go

	//The delete function is used to delete a keyValue out of our map
	delete(colors, "white")
	fmt.Println(colors) //prints: map[green:#4bf745 red:#ff0000]

	printMap(colors)
}

/*
We take a map and interate over all the keyValue paris

Remember that map is a reference type so you dont need to worry about passing
a pointer when you want to update the varibale of type map in a function
*/
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex) //You can imagine color and hex variables as key and value

		//Same rule applies here where you can use undersocre to replace the key "color" if you are not using the key
	}
}
