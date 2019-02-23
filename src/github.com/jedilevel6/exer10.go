package jedilevel6

import (
	"fmt"
)

func VariableClosure() {
	x := 12

	x++
	{
		y := 100
		x++
		y++
		fmt.Println("Value of x is ", x)
		fmt.Println("Value of y is ", y)
	}

	x++
	fmt.Println("Value of x is ", x)
	// below statement throws y undefined error since scope of y is limited to above block
	//fmt.Println("Value of y is ", y)
}
