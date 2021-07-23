package main

import (
	"fmt"
)

func main() {
	/*
	Answer = 24
	Answer = -8
	*/
	var num1 int = 28
	var num2 int = 8
	/*Answer = -8.000000*/
	//answer:= num1 - float64(num2)
	answer2:= num1%num2
	fmt.Printf("Answer = %d", answer2)
}
