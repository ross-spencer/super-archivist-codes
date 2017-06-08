// 1. we always start by declaring a package
// a package is a logical grouping of functions
// when we declarer package main however, we're saying this
// is a standalone application
package main

// 2. you may have guessed by now we can annotate code using forward
// slashes, these are called comments. Comments are good practice
// they can also be used to disable pieces of code when developing

// 3. We should import some ready made packages to help us with our
// code. We don't need to write everything from scratch this pay
import (
	"fmt"
	"math/rand"
	"time"
)

// 4. we then use functions to help split code up. Functions in general
// should do ONE thing WELL. This helps with testing in future. But as
// we develop code, we can go pretty easy on this.

// 5. A simple function is declared as follows
// this function takes no parameters, and has no return value
// it prints two versions of time...
func displayTime() {

	fmt.Println("we call displaytime() function...")

	t := time.Now()

	// The variable t, already contains some date information...
	fmt.Println(t)

	// N.B. The time package is strange. The string below 2006-01-02 15:04:05
	// describes the pattern we want to display our datetime, not the datetime
	// we will display...
	fmt.Println(t.Format("2006-01-02 15:04:05"))

	// try swapping the year around for example
	fmt.Println(t.Format("02-01-2006 15:04:05"))

	// try just displaying the year
	fmt.Println(t.Format("2006"))
}

// 6. A more complex function might take an input
// a parameter is just a variable, so the two inputs here are
// a, and b. Their datatype is int which stands for integer
func addAndPrintNumbers(a int, b int) {

	fmt.Println("we call addAndPrintNumbers() function...")
	fmt.Println(a + b)
}

func addNumbers(a int, b int) int {
	fmt.Println("we call addNumbers() function...")
	fmt.Println("Random number 1:", a)
	fmt.Println("Random number 2:", b)
	return a + b
}

// 7. And finally, a more complex function might return a value to
// the piece of code that called it...
func getRandom() int {

	fmt.Println("we call getRandom() function...")

	// generate a random number
	nanoseconds := time.Now().UnixNano()
	rand.Seed(nanoseconds) // Try changing this number!

	// Return a random number between one and 100
	return rand.Intn(100)
}

// 8. To run all our nice functions, we need a 'main' function in
// our code. It is like the beginning of a book - the entry point for the
// program. Telling it where to start and what to do.
func main() {
	displayTime()
	addAndPrintNumbers(10, 23)
	addAndPrintNumbers(3, 30)

	// while we return a value from getRandom above, we still need to do
	// something with it, else it will do nothing... 
	getRandom()

	// So let's assign it to another variable to print
	aynRand := getRandom()
	fmt.Println(aynRand)
}
