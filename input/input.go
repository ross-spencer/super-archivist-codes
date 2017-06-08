package main 

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func interactive() {

	// 2. Create a new READER to read input from the console
	reader := bufio.NewReader(os.Stdin)

	// 3. Format our text a little to present it to our users...
	fmt.Println("Say hello to all your friends. Type Q to exit.")
	fmt.Print("Enter text: ")

	// 4. Next read our input...
	text, _ := reader.ReadString('\n')

	// 5. We haven't covered this yet, but we clean the string using
	// TrimSpace to remove excess spaces in the input... if you want 
	// to see the spaces, do a Println above this comment... 
	text = strings.TrimSpace(text)

	// 6. Now we want to see if someone is trying to exit, or if we have
	// a name to say hello to...
	if text == "Q" {
		// 7. We can always instruct a program to exit for us...
		os.Exit(0)
	} else {
		fmt.Println("Hello!", text)

		// 8. The final part of this code that is quite important is that the
		// function concludes by calling itself. This helps us to create an 
		// infinite loop that we can control. If the user presses Q and enter 
		// they exit, else we ask the program to repeat the input instructions
		// yhis technique is called 'recursion'
		interactive()
	}
}

func main() {

	// 1. Another nice and simple one, but quite powerful...
	// Check out the interactive function:
	interactive()

}