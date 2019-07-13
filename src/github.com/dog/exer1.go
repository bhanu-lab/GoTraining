// Package dog does have a lot of dog related functions
package dog

// Years ... years calculate total number of dog years by calculating
// human years multiplied by 7
func Years(humanYears int64) int64 {
	dogYears := humanYears * 7
	//fmt.Println("Total dog years is", dogYears)
	return dogYears
}

// YearsTwo ... yearstwo calculates dog years by adding 7 to the times of
// human years
func YearsTwo(humanYears int) int {
	dogYears := 0
	for i := 0; i < humanYears; i++ {
		dogYears += 7
	}
	return dogYears
}

//go doc [package name] is the command used to check documentation on package level
//go doc [package name] [function] is the command used to check documentation on function
//godoc -http=:[portnumber] is the command used to run godoc on local host along with our third party package dicumentation
