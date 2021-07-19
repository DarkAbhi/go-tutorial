package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/**
	Please type something: Hello Dude!
	Typed = "Hello Dude!"
	 */
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please type something: ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("Typed = %q", input)
}
