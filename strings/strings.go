package main 

import (
	"fmt"
	"log"
)

func main() {

	// 1. A lot of the work we do in programming is about working with
	// strings. Lines of text we want to convert into something else.
	// OR other variables that we want to turn into strings. 

	// 2. Simple string output, 
	fmt.Println("Our first string")

	// 3. If we don't use Println (Print Line) we need to add a newline
	// character...
	fmt.Print("Our second string\n")

	// 4. Print Line and Print are equivalent in the log function too
	log.Println("Our first log string")
	log.Print("Our second log string\n")

	// 5. A simple way to output a string of different data types is to
	// use commas 
	fmt.Println("One", 1, true)

	// 6. For two strings we can use the plus sign
	fmt.Println("One " + "two")

	// 7. but for more complicated formatting we need to use another method
	// Printf is described here https://golang.org/pkg/fmt/#Printf
	// the values after the % character are flags that instruct the program
	// how to interpret the value we're outputting... https://golang.org/pkg/fmt/
	fmt.Printf("One %d %t and then some other words.", 1, true)

	// There is a lot we will want to do with strings over time, so keep an
	// eye out for different use-cases.
}