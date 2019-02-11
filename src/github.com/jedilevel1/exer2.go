package main

import "fmt"

func PackLevelVariables() {

	var (
		x int
		y string
		z bool
	)

	fmt.Println(x, y, z)

	fmt.Println("X values is ", x)
	fmt.Println("Y values is ", y)
	fmt.Println("Z values is ", z)

}
