package jedilevel10

import (
	"fmt"
	"sync"
)

func testRoutinesIncAndDec() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		for i := 1; i <= 100; i++ {
			fmt.Println(i)
		}
		fmt.Println("Incrementer is completed")
		wg.Done()
	}()

	go func() {
		for i := 100; i >= 1; i-- {
			fmt.Println(i)
		}
		fmt.Println("Decrementer is completed")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Incrementer and Decrementer is completed")
}
