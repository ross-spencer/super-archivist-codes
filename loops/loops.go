package main 

import (
	"fmt"
)

func printHello(i int) {
	fmt.Println("Hello", i)
}

// This might be the first time we see a 'byte' data type
// this is just what it says on the tin, and can be converted to
// other data types if necessary, here we make it into a character to
// print our letters one by one...
func printGoodbye(i int, s byte) {
	fmt.Println("Goodbye", i, string(s))		// string() allows us to convert (cast) the datatype
}															// without string() here it would print a numberical value

func main() {

	// 1. Loops are cool. They help us to repeat the same opeation
	// over and over again. 

	// 2. We have two main ways we can make a loop... each begins 
	// with the word for...
	// the structure is: initial condition; limit (when to stop); instructions for the counter
	// here, ++ means add one each loop, -- would remove one each loop...
	for x := 0; x < 5; x++ {
		printHello(x)	// x is a counter, we can see it get bigger
	}

	// 3. If we want to loop over some known information, for 
	// example, a string... we can do that this way...
	var word = "Supercalifragilisticexpialidocious"
	for x := range(word) {
		printGoodbye(x, word[x])	// we can even use our count as an 'index' to print the letter
	}

}