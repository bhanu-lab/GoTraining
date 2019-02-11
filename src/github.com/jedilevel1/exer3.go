package main

import "fmt"

var (
	x = 42
	y = "James Bond"
	z = true
)

func FormatValuesAndPrint() {
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
}
