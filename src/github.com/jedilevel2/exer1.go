package main

import (
	"fmt"
)

// DecimalNumberInDiffFormats ... Prints decimal, binary, hexadecimal formats of a given number
func DecimalNumberInDiffFormats(x int) {

	// %b is for binary format
	// %x is for hexa decimal format
	// %d is for decimal format
	fmt.Printf("%d \t %b \t %x", x, x, x)
}
