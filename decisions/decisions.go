package main  

import (
	"fmt"
)

func main() {

	// 1. We're going to look at a handful of combinations of IF
	// statement. They help us to route code through various functions.

	// 2. Equivalence
	var a = 1
	var b = 1

	if a == b {
		fmt.Println("these variables have the same information!", a, b)
	} else {
		fmt.Println("these variables hold different information", a, b)
	}

	// 3. Testing variables are different (just the inverse)
	var d = "hello"
	var e = "goodbye"

	if d != e {
		fmt.Println("these variables are different", d, e)
	} else {
		fmt.Println("these variables are the same", d, e)
	}

	// 4. And we can ask more complex questions of our data by building
	// on these principles...

	// 5. AND using the syntax: &&
	if a == b && d == e {
		fmt.Println("a equals b AND d equals e", a, b, d, e)
	} else {
		fmt.Println("a doesn't equal b AND d doesn't equal e", a, b, d, e)
	}

	// 6. OR using the syntax ||
	if a == b || d != e {
		fmt.Println("either a equals b OR d doesn't equal e", a, b, d, e)
	} else {
		fmt.Println("neither a equals b OR d doesn't equal e", a, b, d, e)
	}

	// 7. And we can build if statements using an ELSE IF keyword
	// it's like a backup if the first IF fails
	if a != b {
		fmt.Println("a doesn't equal b")
	} else if d != e {
		fmt.Println("d doesn't equal e")		
	} else {
		fmt.Println("neither of our tests worked")
	}

	// A good exercise using loops and IF statements is FIZZBUZZ:
	// Write a program that prints the numbers from 1 to 100. But for 
	// multiples of three print "Fizz" instead of the number and for 
	// the multiples of five print "Buzz" 
}