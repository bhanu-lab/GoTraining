package jedilevel3

import (
	"fmt"
)

// PrintOneToThousandUsingForLoop ... prints all the numbers from 1 to 10000
func PrintOneToThousandUsingForLoop() {
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
}
