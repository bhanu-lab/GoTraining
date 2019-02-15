package jedilevel3

import (
	"fmt"
)

// DetermineIfElseBlocksUsage ... determines if else blocks usage
func DetermineIfElseBlocksUsage(x int) {

	if x < 10 {
		fmt.Println("X is less that 10")
	} else if x == 10 {
		fmt.Println("X is equal to 10")
	} else {
		fmt.Println("X is greater than 10")
	}
}
