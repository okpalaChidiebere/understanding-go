/*
How is the go file code organized?

At the very top, the first line of code we always have the package declaration.

Secondly, we list out any package dependencies that we might need to install
into the working file. These could be imported from the go Standard library
or custom reusable packages.

Lastly the body of the file which is where we write a bunch of logic to do
stuff we like! eg variable declarations, etc

*/

/*
- Package declaration: The keyword "main" means all code in this file will be generated
an executable file for your code to run. Any other keyword eg package claculator means that the code
is resuable and no exe code will be generated for it
*/
package main

/*We use import statement to gain access to another package inside the one that we are currently authoring
We can import code from go standard libraries https://golang.org/pkg/ or import code from
other engineer programmers as well*/
import "fmt"

/*
After imports you start to write the logic for your code.

NOTE: Every Executable file must have this main function */
func main() {
	//call the printLine function
	fmt.Println("Hi there")
}

/*
How we run our go project?
We run the file using the command line interface.

We can instantly run our code or

we can build our code and run it in the future at some point ourself. We will use
"go build <MAIN_FILE>.go" to generate an executable that we can run later with "./<MAIN_FILE_NAME>"
*/
