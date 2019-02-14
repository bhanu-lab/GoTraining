package main

import (
	"fmt"
)

const (
	a = 10
	b = 32.34334
	c = "Billa Bong"
)

const (
	d int     = 21
	e float64 = 343.34334
	f string  = "Bead Billa"
)

//PrintConstValues ... prints constant values assigned
func PrintConstValues() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}

// PrintConstTypeValues ... prints constant typed values
func PrintConstTypeValues() {
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)
}

func main() {
	PrintConstValues()
	PrintConstTypeValues()
}
