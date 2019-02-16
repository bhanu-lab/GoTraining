package jedilevel4

import (
	"fmt"
)

// ArrayIteration ... added values and iterating through them using range
func ArrayIteration() {
	var num = [5]int{11, 12, 13, 14, 15}
	/*sam := 10
	for i, _ := range num {
		num[i] = sam
		sam += 1
	}*/

	for i, v := range num {
		fmt.Println("value is ", i, v)
	}

	fmt.Printf("data type is %T", num)
}
