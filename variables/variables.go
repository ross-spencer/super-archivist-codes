package main 

import (
	"fmt"
	"reflect"
)

func main() {

	// 1. A variable can be created a number of ways
	// := asks the compiler to figure out what datatype it is
	// here first will be an integer (int) type
	first := 33

	// 2. This is the equivalent of the above, but is more verbose
	var second = 33

	// 3. Even more verbose again we specify the exact data type for
	// ourselves... 
	var third int
	third = 33

	// 4. We can use a package called reflect to see the different data types
	fmt.Println(first, second, third)
	fmt.Println(reflect.TypeOf(first), reflect.TypeOf(second), reflect.TypeOf(third))

	// 5. Other data types that are useful are, bool, and string
	fourth := true
	fifth := "some text"

	fmt.Println(fourth, fifth)
	fmt.Println(reflect.TypeOf(fourth), reflect.TypeOf(fifth))

	// Other golang data types are listed here: https://tour.golang.org/basics/11
	// Our choices around datatype outside of those above, are either for more
	// special uses, or to optimize our code, e.g. using less memory for some data. 
}