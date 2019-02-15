package main

import "fmt"

func main() {

	born := 1989

	for born <= 2089 {
		fmt.Printf("Year %v \n", born)
		born++
	}
}
