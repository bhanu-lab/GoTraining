package jedilevel4

import (
	"fmt"
)

// SlicingSlice ... performing different slicing operations
func SlicingSlice() {

	ele := []int{41, 42, 43, 44, 45, 46, 47, 48, 49, 50}

	fmt.Println(ele[1:7])
	ele = append(ele, 51)
	fmt.Println(ele[7:])
	fmt.Println(ele[4:9])
	fmt.Println(ele[3:8])
}
