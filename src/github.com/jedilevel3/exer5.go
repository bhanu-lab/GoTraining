package jedilevel3

import (
	"fmt"
)

// PrintsRemainderFrom10To100With4 ... prints remainder when divided by 4
func PrintsRemainderFrom10To100With4() {
	num := 10

	for num <= 100 {
		rem := num % 4
		fmt.Println("Remainader for", num, "is", rem)
		num++
	}
}
