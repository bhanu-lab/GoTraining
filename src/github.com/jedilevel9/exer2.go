package jedilevel9

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func PrintNumbersConcurrently() {
	fmt.Println("testing ...")

	wg.Add(1)

	go PrintFromOneToHundred()
	go PrintFromHundredToOne()

	wg.Wait()
}

func PrintFromOneToHundred() {

	for i := 1; i <= 100; i++ {
		fmt.Println("Number is ", i)
	}
}

func PrintFromHundredToOne() {

	for i := 100; i >= 1; i-- {
		fmt.Println("Number is ", i)
	}

	wg.Done()
}
