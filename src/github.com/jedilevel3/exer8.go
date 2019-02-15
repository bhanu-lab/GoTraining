package jedilevel3

import (
	"fmt"
)

// UseSwitchCaseWithoutCondition ... uses switch without any expression in it
func UseSwitchCaseWithoutCondition() {

	switch {

	case 2 == 2:
		fmt.Println("Two is always equal to Two")
		fallthrough
	case 3 == 2:
		fmt.Println("Second switch casee")
		fallthrough
	case 4 == 4:
		fmt.Println("Four is always equal to Four")

	}
}
