package jedilevel8

import (
	"fmt"
	"sort"
)

// SortUsers ... sorts slices with standard package sort
func SortUsers() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	sort.Ints(xi)

	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)

	fmt.Println(xs)
}
