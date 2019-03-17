// Package dog does have a lot of dog related functions
package dog

import (
	"fmt"
)

// Years ... years calculate total number of dog years by calculating
// human years multiplied by 7
func Years(humanYears int64) int64 {
	dogYears := humanYears * 7
	fmt.Println("Total dog years is", dogYears)
	return dogYears
}

//go doc [package name] is the command used to check documentation on package level
//go doc [package name] [function] is the command used to check documentation on function
//godoc -http=:[portnumber] is the command used to run godoc on local host along with our third party package dicumentation
