package jedilevel4

import "fmt"

// InitializeAMap .. create a map with default values
func InitializeAMap() {
	x := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"money_pennymiss": []string{"James Bond", "Literature", "ComputerScience"},
		"no_dr":           []string{"BeingEvil", "IceCream", "Sunsets"},
	}

	for i, slic := range x {
		fmt.Println("Thisis the record for ", i)
		for k, val := range slic {
			fmt.Println("\t", k, val)
		}
	}
}
