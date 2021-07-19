package main

import "fmt"

func main() {
	var x string = fmt.Sprintf("Hello %T %v", 10,10)
	println(x)
	/**
	Hello int 10
	 */
}
