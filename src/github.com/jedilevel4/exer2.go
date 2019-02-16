package jedilevel4

import "fmt"

func IterateSlices() {
	var numbers []int
	numbers = append(numbers, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)

	for i, v := range numbers {
		fmt.Printf("Value at position %v is %v\n", i, v)
	}

	fmt.Printf("type os slice is %T", numbers)
}
