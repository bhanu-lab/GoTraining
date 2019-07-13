package main

import (
	"fmt"
)

// ShiftBits ... prints digit in binary and hex left shifts bits by 1 and re prints decimal, binary, hexadecimal
func ShiftBits(dig int) {

	fmt.Printf("%d\t%b\t%x\n", dig, dig, dig)

	shiftedDig := dig << 1

	fmt.Printf("%d\t%b\t%x", shiftedDig, shiftedDig, shiftedDig)
}
