package jedilevel3

import (
	"fmt"
)

// PrintUnicodeCharactersFromAToZThrice ... prints A to Z characters unicode thrice
func PrintUnicodeCharactersFromAToZThrice() {

	for char := 'A'; char <= 'Z'; char++ {
		for count := 0; count < 3; count++ {
			fmt.Printf("%#U", char)
		}
		fmt.Println("")
	}
}
