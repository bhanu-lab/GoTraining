package jedilevel6

import (
	"fmt"
)

func InvokeFuntions() {

	foo()
	moo("Thuglaq")
	sum := count(1, 2, 3, 4, 5, 6, 7, 8, 9)
	sum, aboveZero := countAndCheckIfGreaterThanZero(1, 2, 4, 4, 5, 6, 7, 8, 9)

	fmt.Println("Total value is ", sum, aboveZero)

}

func foo() {
	fmt.Println("function with out any parameters")
}

func moo(name string) {
	fmt.Println("parameters received to moo", name)
}

func count(x ...int) int {

	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum
}

func countAndCheckIfGreaterThanZero(x ...int) (int, bool) {
	sum := count(x...)
	sumGreaterThanZero := false
	if sum > 0 {
		sumGreaterThanZero = true
	}

	return sum, sumGreaterThanZero
}
