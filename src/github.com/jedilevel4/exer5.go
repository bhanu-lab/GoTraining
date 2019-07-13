package jedilevel4

import "fmt"

// DeleteSlices ... delete and assign to a different slice
func DeleteSlices() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y := x[:3]

	y = append(y, y[6:]...)

	fmt.Println(y)

}
