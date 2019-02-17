package jedilevel4

import (
	"fmt"
)

// MultiDimensionalSlice ... creating a miltidimensional slice
func MultiDimensionalSlice() {
	values := [][]string{}

	row1 := []string{"James", "Bond", "Shaken, Stirred"}
	row2 := []string{"Miss", "Money penny", "Hello, James"}

	values = append(values, row1)
	values = append(values, row2)

	for i, v := range values {
		fmt.Println(i, "has value of ", v)

		for j, data := range v {
			fmt.Println(j, "has data of", data)
		}
	}
}
