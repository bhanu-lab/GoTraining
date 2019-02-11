package main

import (
	"fmt"
)

func main() {
	var n int //declaration of variables
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			n, _ = fmt.Println(i, "is the even number") //Println -> accepts variadic parameters which means you can pass any number of parameters
			//interface{} -> every value in go is of type interface{}
			fmt.Println(n, "number of bytes printed")
		} else {
			fmt.Println(i, "is odd number")
		}
	}
}
