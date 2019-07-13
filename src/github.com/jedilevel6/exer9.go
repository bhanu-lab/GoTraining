package jedilevel6

import "fmt"

// CallingBackAFunction ... creates an anonymous function and passed it to another function
func CallingBackAFunction(x ...int) int {
	f := func(x ...int) int {
		var total int
		for _, v := range x {
			total = total + v
		}

		return total
	}

	return TestCallBackFunc(f, x...)
}

// TestCallBackFunc ... callback function
func TestCallBackFunc(x func(z ...int) int, i ...int) int {
	total := x(i...)

	fmt.Println("Total sum is ", total)
	return total
}
