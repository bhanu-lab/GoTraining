package jedilevel4

import (
	"fmt"
)

// AddToMap ... Adding values to map
func AddToMap() {
	x := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"money_pennymiss": []string{"James Bond", "Literature", "ComputerScience"},
		"no_dr":           []string{"BeingEvil", "IceCream", "Sunsets"},
	}

	x["dummy"] = []string{"Horror", "comedy", "scifi"}

	for i, slic := range x {
		fmt.Println("Thisis the record for ", i)
		for k, val := range slic {
			fmt.Println("\t", k, val)
		}
	}
}
