package jedilevel6

import (
	"fmt"
)

func UsingDeferFUnction() {

	defer meow()
	mowmow()
}

func meow() {
	fmt.Println("foo is deffered and executed at the end of the enclosing function")
}

func mowmow() {
	fmt.Println("hehehehe .... I will be invoked earlier than meow even though foo is infront of me")
}
