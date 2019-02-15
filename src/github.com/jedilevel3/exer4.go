package main

import (
	"fmt"
)

func main() {

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
