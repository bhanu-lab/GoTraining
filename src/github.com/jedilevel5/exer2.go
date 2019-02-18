package jedilevel5

import (
	"fmt"
)

// AddingStructToMap ... adding struct to a map as value
func AddingStructToMap() {
	x := example{
		firstName:   "Test",
		lastName:    "team",
		favIceCream: []string{"chocolate", "strawberry", "hazzlenut"},
	}

	var map1 map[string]example

	map1 = make(map[string]example)

	map1[x.lastName] = x

	y := example{
		firstName:   "test1",
		lastName:    "team1",
		favIceCream: []string{"chocolate1", "banana", "custard"},
	}

	map1[y.lastName] = y

	for k, v := range map1 {
		fmt.Println(k, v)

		for _, val := range v.favIceCream {
			fmt.Println(val)
		}
	}
}
