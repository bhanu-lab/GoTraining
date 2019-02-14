package main

import (
	"fmt"
)

// UseComparisonOperators ... used to print results of all comparision operators between two values
func UseComparisonOperators(a int, b int) {

	z := a == b
	y := a <= b
	x := a >= b
	w := a != b
	v := a < b
	u := a > b

	fmt.Println(" a == b value is ", z)
	fmt.Println(" a <= b value is ", y)
	fmt.Println(" a >= b value is ", x)
	fmt.Println(" a != b value is ", w)
	fmt.Println(" a < b value is ", v)
	fmt.Println(" a > b value is ", u)
}

func main() {
	UseComparisonOperators(21, 32)
}
