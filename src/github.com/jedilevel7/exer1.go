package jedilevel7

import (
	"fmt"
)

func GetVariableAddress() {
	i := 26
	fmt.Println("Printing the value", i)
	fmt.Println("Printing the address of value", &i)
	fmt.Printf("%T\n", &i)
}
