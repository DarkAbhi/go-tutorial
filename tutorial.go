package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	/**
	Year you were born? 2000
	You are 21 years old
	*/
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Year you were born? ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You are %d years old", 2021-input)
}
