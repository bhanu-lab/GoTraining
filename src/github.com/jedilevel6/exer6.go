package jedilevel6

import (
	"fmt"
)

// AnonymousFunction defining an anonymous function
func AnonymousFunction() {
	func(x int) {

		if x%2 == 0 {
			fmt.Println("Given number is even")
		} else {
			fmt.Println("Given number is odd")
		}
	}(5)
}
