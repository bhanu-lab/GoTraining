package main

import "fmt"

type raspi int

var jeffa raspi

func CreateDataTypeAndAssignValue() {

	fmt.Println(jeffa)
	fmt.Printf("%T\n", jeffa)
	jeffa = 42
	fmt.Println(jeffa)

}
