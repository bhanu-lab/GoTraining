package main

import "fmt"

const (
	year = 2019 + iota
	nextYear
	afterTwoYears
	afterThreeYears
)

func UsingiotaType() {

	fmt.Println("Current year is ", year)
	fmt.Println("Next year is ", nextYear)
	fmt.Println("After two years is ", afterTwoYears)
	fmt.Println("After three years is ", afterThreeYears)
}

func main() {
	UsingiotaType()
}
