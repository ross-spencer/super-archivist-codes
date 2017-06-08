package main 

import (
	"fmt"
)

func main() {

	// 1. We'll keep this one simple because it can get quite complex. 
	// Maps and slices help us to store groups of information 

	// First, a slice (this is similar to what is called an array in other languages)
	var fruit = [...]string{"Apples","Bananas","Pears","Peaches"}

	// 2. The ... syntac asks the compiler to count the values for itself
	// we could also just type 4

	// 3. we can loop through and print there...

	for x := range(fruit) {
		fmt.Println("Fruit", x, ":", fruit[x])	// x is an index that allows us to access the values
	}					

	// 4. Indexes always start at zero, so the first element, is the zeroth
	// zero-based indexes can be tricky, but we learn more as we use them
	fmt.Println("Element zero:",fruit[0])
	fmt.Println("Element three:",fruit[3])

	// 5. A Map is similar to a slice
	fruitmap := make(map[string]int)

	// 6. lets keep track of how much fruit we have... 
	fruitmap["Apples"] = 1
	fruitmap["Bananas"] = 2
	fruitmap["Pears"] = 3
	fruitmap["Peaches"] = 4

	// N.B. Run the code a few times, when they output, they
	// will not always be in the same order every time, this is
	// a design decision by Golang... We have to do the ordering.
	for x, y := range(fruitmap) {
		fmt.Println(x, y)
	}

	// 7/ How many apples do we have?
	fmt.Println("We have", fruitmap["Apples"], "Apples")
}