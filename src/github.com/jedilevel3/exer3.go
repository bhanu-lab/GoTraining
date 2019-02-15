package jedilevel3

import "fmt"

// PrintAllYearsAlive ... prints all the alive
func PrintAllYearsAlive() {

	born := 1989

	for born <= 2089 {
		fmt.Printf("Year %v \n", born)
		born++
	}
}
