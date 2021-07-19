package main

import "fmt"

func main() {
	/**
	Value true
	*/
	fmt.Printf("\nValue %t", true)

	/**
	Bytes 10000000000
	*/
	fmt.Printf("\nBytes %b", 1024)

	/**
	Octal 6553
	*/
	fmt.Printf("\nOctal %o", 3435)

	/**
	Decimal d6b
	Decimal D6B
	%x for lower case, %X for upper case
	*/
	fmt.Printf("\nDecimal %X", 3435)

	/**
	Exponent 2.373562e+00
	*/
	fmt.Printf("\nExponent %e", 2.373561516516)

	/**
	Decimal no exponent 2.373562
	*/
	fmt.Printf("\nDecimal no exponent %F", 2.373561516516)

	/**
	Large decimal 2.373562
	*/
	fmt.Printf("\nDecimal no exponent %g", 2.373561516516)

	/**
	String Abhi
	Use %q to enclose in double quotes
	*/
	fmt.Printf("\nString %q","Abhi")

	/**
	String    "Abhi" is cool
	Use negative number to left justify
	String "Abhi"    is cool
	*/
	fmt.Printf("\nString %-9q is cool","Abhi")

	/**
	Decimal points 2.47
	Decimal points 2
	*/
	fmt.Printf("\nDecimal points %.2F",2.4654651651)
	fmt.Printf("\nDecimal points %.F",2.4654651651)

	/**
	Padding with digit
	Digit 0000045
	*/
	fmt.Printf("\nDigit \t%07d",45)
}
