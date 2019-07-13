package jedilevel3

import (
	"fmt"
)

// PrintYearsAliveWithForLoop ... prints all years alive using for true statement
func PrintYearsAliveWithForLoop() {

	born := 1989

	for {
		if born <= 2089 {
			fmt.Printf("Year is %v \n", born)
		} else {
			break
		}
		born++
	}
}
